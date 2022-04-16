package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/timvosch/gospot/pkg/pb"
)

var (
	ErrLoginFailed         = errors.New("auth failure")
	ErrKeyNegotationFailed = errors.New("key negotiation failed")
	ErrConnectionFailed    = errors.New("connection failed")
)

const (
	AP_ENDPOINT         = "http://apresolve.spotify.com:80"
	AP_FALLBACK         = "ap.spotify.com:443"
	PACKET_LOGIN        = 0xab
	PACKET_AP_WELCOME   = 0xac
	PACKET_AUTH_FAILURE = 0xad
)

// Connection ...
type Connection struct {
	ReusableCredentials *pb.LoginCredentials
	credentials         *pb.LoginCredentials
	deviceID            string
	stream              *secureStream
}

func Connect(cred *pb.LoginCredentials, deviceID string) (*Connection, error) {
	conn := &Connection{
		credentials: cred,
		deviceID:    deviceID,
	}

	// Find a spotify access point
	ap := getAP()
	log.Printf("[Connection] Connecting to %s...", ap)
	sock, err := net.Dial("tcp", ap)
	if err != nil {
		return nil, ErrConnectionFailed
	}

	// Negotiate keys
	insecure := newInsecureStream(sock)
	keys, err := insecure.NegotiateKeys()
	if err != nil {
		sock.Close()
		log.Printf("[Connection] Key negotiation failed: %s", err)
		return nil, ErrKeyNegotationFailed
	}

	// Create secure stream
	conn.stream = newSecureStream(sock, keys)

	// Login
	log.Printf("[Connection] Logging in for %s\n", *cred.Username)
	if err := conn.login(); err != nil {
		sock.Close()
		log.Printf("[Connection] Login failed: %s", err)
		return nil, fmt.Errorf("%w: %s", ErrLoginFailed, err)
	}

	log.Println("[Connection] Login successful")
	return conn, nil
}

func (c *Connection) Close() {
	c.stream.Close()
}

func (c *Connection) login() error {
	// TODO: check if secureStream and credentials are set

	// Create and send login request
	clientLogin := &pb.ClientResponseEncrypted{
		LoginCredentials: c.credentials,
		SystemInfo: &pb.SystemInfo{
			CpuFamily:               pb.CpuFamily_CPU_UNKNOWN.Enum(),
			Os:                      pb.Os_OS_UNKNOWN.Enum(),
			SystemInformationString: proto.String("gospot_01_01"),
			DeviceId:                proto.String(c.deviceID),
		},
		VersionString: proto.String("gospot_01_01"),
	}
	clientLoginRaw, err := proto.Marshal(clientLogin)
	if err != nil {
		return fmt.Errorf("failed to marshal login packet: %s", err)
	}
	if err := c.stream.Send(PACKET_LOGIN, clientLoginRaw); err != nil {
		return fmt.Errorf("failed to send login packet: %s", err)
	}

	// Get and parse response
	cmd, serverLoginRaw, err := c.stream.Receive()
	if err != nil {
		return fmt.Errorf("failed to receive login response: %v", err)
	}
	// failed
	if cmd == PACKET_AUTH_FAILURE {
		return getLoginFailedReason(serverLoginRaw)
	} else if cmd != PACKET_AP_WELCOME {
		return fmt.Errorf("unexpected packet: %d", cmd)
	}
	// success
	var serverLogin pb.APWelcome
	if err := proto.Unmarshal(serverLoginRaw, &serverLogin); err != nil {
		return fmt.Errorf("failed to unmarshal login response: %v", err)
	}

	// Create new credentials
	c.ReusableCredentials = &pb.LoginCredentials{
		Username: c.credentials.Username,
		Typ:      serverLogin.ReusableAuthCredentialsType.Enum(),
		AuthData: serverLogin.ReusableAuthCredentials,
	}

	return nil
}

func getLoginFailedReason(data []byte) error {
	var pkt pb.APLoginFailed
	if err := proto.Unmarshal(data, &pkt); err != nil {
		return fmt.Errorf("failed to unmarshal login failed packet : %v", err)
	}

	return fmt.Errorf("spotify reported: %s", *pkt.ErrorDescription)
}

// apListResponse ...
type apListResponse struct {
	Items []string `json:"ap_list"`
}

func getAP() string {
	res, err := http.Get(AP_ENDPOINT)
	if err != nil {
		return AP_FALLBACK
	}

	var apList apListResponse
	if err := json.NewDecoder(res.Body).Decode(&apList); err != nil {
		return AP_FALLBACK
	}

	return apList.Items[rand.Intn(len(apList.Items))]
}
