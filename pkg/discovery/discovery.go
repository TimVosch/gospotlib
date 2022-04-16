package discovery

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"

	"github.com/grandcat/zeroconf"
	"github.com/timvosch/gospot/internal/crypto"
	"github.com/timvosch/gospot/pkg/pb"
)

var (
	ErrShutdown      = errors.New("discovery shutdown")
	ErrBlobMalformed = errors.New("blob malformed")
)

// Discovery ...
type Discovery struct {
	deviceName  string
	deviceID    string
	keys        crypto.PrivateKeys
	pubKey      string
	credentialC chan *pb.LoginCredentials
	done        chan struct{}
}

func New(name string) *Discovery {
	keys := crypto.GenerateKeys()

	d := &Discovery{
		deviceName:  name,
		deviceID:    generateDeviceID(name),
		keys:        keys,
		pubKey:      base64.StdEncoding.EncodeToString(keys.PubKey()),
		credentialC: make(chan *pb.LoginCredentials, 1),
	}

	return d
}

func (d *Discovery) DeviceID() string {
	return d.deviceID
}

func (d *Discovery) Credentials() <-chan *pb.LoginCredentials {
	return d.credentialC
}

func (d *Discovery) Start() error {
	errC := make(chan error, 1)
	d.done = make(chan struct{}, 1)
	defer close(d.done)
	defer log.Println("[Discovery] Stopped")

	// Find free socket
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	defer l.Close()
	log.Printf("[Discovery] Bound to %s\n", l.Addr().String())

	// Start HTTP Server
	httpSrv := &http.Server{
		Handler: d.createHandler(d.credentialC),
	}
	go func() {
		if err := httpSrv.Serve(l); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
		d.Shutdown()
	}()
	defer httpSrv.Close()
	log.Println("[Discovery] HTTP server online")

	// Start zeroconf server
	port := l.Addr().(*net.TCPAddr).Port
	instance := fmt.Sprintf("%s-%d", d.deviceName, rand.Intn(1000))
	zcSrv, err := zeroconf.Register(
		instance,
		"_spotify-connect._tcp",
		"local.",
		port,
		[]string{"CPath=/"},
		nil,
	)
	if err != nil {
		return err
	}
	defer zcSrv.Shutdown()
	log.Println("[Discovery] Advertising on network")

	// Wait for shutdown or error
	select {
	case <-d.done:
		return ErrShutdown
	case err := <-errC:
		return err
	}
}

func (d *Discovery) Shutdown() {
	log.Println("[Discovery] Shutting down...")
	close(d.done)
}

func (d *Discovery) createHandler(c chan *pb.LoginCredentials) http.Handler {
	// getInfoResponse ...
	type getInfoResponse struct {
		AccountReq       string `json:"accountReq"`
		ActiveUser       string `json:"activeUser"`
		BrandDisplayName string `json:"brandDisplayName"`
		DeviceID         string `json:"deviceID"`
		DeviceType       string `json:"deviceType"`
		GroupStatus      string `json:"groupStatus"`
		LibraryVersion   string `json:"libraryVersion"`
		ModelDisplayName string `json:"modelDisplayName"`
		PublicKey        string `json:"publicKey"`
		RemoteName       string `json:"remoteName"`
		ResolverVersion  string `json:"resolverVersion"`
		SpotifyError     int    `json:"spotifyError"`
		Status           int    `json:"status"`
		StatusString     string `json:"statusString"`
		Version          string `json:"version"`
		VoiceSupport     string `json:"voiceSupport"`
	}
	// addUserResponse ...
	type addUserResponse struct {
		Status       int    `json:"status"`
		StatusString string `json:"statusString"`
		SpotifyError int    `json:"spotifyError"`
	}
	mw := func(rw http.ResponseWriter, r *http.Request) {
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}
		log.Printf("[Discovery] HTTP Request. Method=%s Action=%s URL=%s\n", r.Method, action, r.URL.String())

		switch action {
		case "getInfo":
			{
				body, err := json.Marshal(getInfoResponse{
					AccountReq:       "PREMIUM",
					ActiveUser:       "",
					BrandDisplayName: "gospot",
					DeviceID:         d.deviceID,
					DeviceType:       "Speaker",
					GroupStatus:      "NONE",
					LibraryVersion:   "0.0.1",
					ModelDisplayName: "gospot",
					PublicKey:        d.pubKey,
					RemoteName:       d.deviceName,
					ResolverVersion:  "0",
					SpotifyError:     0,
					Status:           101,
					StatusString:     "ERROR-OK",
					Version:          "2.7.1",
					VoiceSupport:     "NO",
				})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					log.Printf("[Discovery] JSON marshal failed: %s", err)
					return
				}

				rw.Header().Set("Content-Type", "application/json")
				rw.WriteHeader(http.StatusOK)
				rw.Write(body)
			}
		case "addUser":
			{
				username := r.FormValue("userName")
				blob64 := r.FormValue("blob")
				clientKey := r.FormValue("clientKey")
				log.Printf("[Discovery] Adding user...: username=%s, blob=(len: %d), clientKey=(len: %d)\n", username, len(blob64), len(clientKey))

				// Decrypt blob to credentials
				credentials, err := decryptBlob(blob64, clientKey, &d.keys, d.deviceID, username)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					log.Printf("[Discovery] Decrypt blob failed: %s", err)
					return
				}

				log.Printf("[Discovery] Got authentication credentials for: username=%s\n", username)
				d.credentialC <- credentials

				body, err := json.Marshal(addUserResponse{
					Status:       101,
					StatusString: "OK",
					SpotifyError: 0,
				})
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}

				rw.Header().Set("Content-Type", "application/json")
				rw.WriteHeader(http.StatusOK)
				rw.Write(body)
			}
		default:
			http.Error(rw, "unknown action", http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(mw)
}

func generateDeviceID(name string) string {
	hash := sha1.Sum([]byte(name))
	return hex.EncodeToString(hash[:])
}

func decryptBlob(b64, clientKeyB64 string, keys *crypto.PrivateKeys, deviceID, username string) (*pb.LoginCredentials, error) {
	// Decode base64 strings
	blobEnc, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	clientKey, err := base64.StdEncoding.DecodeString(clientKeyB64)
	if err != nil {
		return nil, err
	}

	// Prepare encryption info (IV, chk, payload)
	iv := blobEnc[:16]
	payloadEnc := blobEnc[16 : len(blobEnc)-20]
	chksum := blobEnc[len(blobEnc)-20:]

	sharedKey := keys.SharedKey(clientKey)
	sharedKeyHash := sha1.Sum(sharedKey)
	baseKey := sharedKeyHash[:16]
	baseKeyH := hmac.New(sha1.New, baseKey)

	// Create keys
	baseKeyH.Write([]byte("checksum"))
	chkKey := baseKeyH.Sum(nil)
	baseKeyH.Reset()

	baseKeyH.Write([]byte("encryption"))
	encKey := baseKeyH.Sum(nil)
	baseKeyH.Reset()

	// Verify checksum
	chkHash := hmac.New(sha1.New, chkKey)
	chkHash.Write(payloadEnc)
	mac := chkHash.Sum(nil)
	if !bytes.Equal(chksum, mac) {
		return nil, errors.New("checksum mismatch")
	}

	// Decode payload
	blk, err := aes.NewCipher(encKey[:16])
	if err != nil {
		return nil, fmt.Errorf("failed to create aes cipher for encrypted blob payload: %s", err)
	}
	stream := cipher.NewCTR(blk, iv)
	stream.XORKeyStream(payloadEnc, payloadEnc)

	// Convert decrypted blob to LoginCredentials
	// Note that the blob was double encrypted, the LoginCredentials conversion will do the last decryption
	cred, err := pb.CredentialsFromBlob(payloadEnc, username, deviceID)
	if err != nil {
		return nil, fmt.Errorf("error creating credentials from decrypted blob: %s", err)
	}

	return cred, nil
}
