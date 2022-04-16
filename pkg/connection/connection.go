package connection

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/timvosch/gospot/pkg/pb"
)

var (
	ErrAuthFailure = errors.New("auth failure")
)

const (
	AP_ENDPOINT = "http://apresolve.spotify.com:80"
	AP_FALLBACK = "ap.spotify.com:443"
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
	sock, err := net.Dial("tcp", ap)
	if err != nil {
		return nil, err
	}

	// Negotiate keys
	insecure := newInsecureStream(sock)
	keys, err := insecure.NegotiateKeys()
	if err != nil {
		sock.Close()
		return nil, err
	}

	// Create secure stream
	conn.stream = newSecureStream(sock, keys)

	// Login
	if err := conn.login(); err != nil {
		sock.Close()
		return nil, err
	}

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
		return err
	}
	if err := c.stream.Send(PacketLogin, clientLoginRaw); err != nil {
		return err
	}

	// Get and parse response
	cmd, serverLoginRaw, err := c.stream.Receive()
	if err != nil {
		return err
	}
	// failed
	if cmd == PacketAuthFailure {
		return ErrAuthFailure
	} else if cmd != PacketAPWelcome {
		return errors.New("unexpected packet")
	}
	// success
	var serverLogin pb.APWelcome
	if err := proto.Unmarshal(serverLoginRaw, &serverLogin); err != nil {
		return err
	}

	// Create new credentials
	c.ReusableCredentials = &pb.LoginCredentials{
		Username: c.credentials.Username,
		Typ:      serverLogin.ReusableAuthCredentialsType.Enum(),
		AuthData: serverLogin.ReusableAuthCredentials,
	}

	return nil
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
