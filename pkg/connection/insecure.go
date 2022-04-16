package connection

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"

	"github.com/timvosch/gospot/internal/crypto"
	"github.com/timvosch/gospot/pkg/pb"
	"google.golang.org/protobuf/proto"
)

// insecureStream ...
type insecureStream struct {
	stream io.ReadWriteCloser
}

func newInsecureStream(stream io.ReadWriteCloser) *insecureStream {
	return &insecureStream{stream}
}

func (c *insecureStream) NegotiateKeys() (*crypto.SharedKeys, error) {
	keys := crypto.GenerateKeys()

	clientHelloRaw := makeHelloMessage(keys.PubKey(), keys.ClientNonce())
	clientHelloPkt, err := c.send([]byte{0x00, 0x04}, clientHelloRaw)
	if err != nil {
		return nil, err
	}
	// get response
	serverHelloRaw, err := c.receive()
	if err != nil {
		return nil, err
	}
	var serverHello pb.APResponseMessage
	if err := proto.Unmarshal(serverHelloRaw[4:], &serverHello); err != nil {
		return nil, err
	}

	remoteKey := serverHello.GetChallenge().GetLoginCryptoChallenge().GetDiffieHellman().GetGs()
	sharedKeys := keys.AddRemoteKey(remoteKey, clientHelloPkt, serverHelloRaw)

	clientRes := &pb.ClientResponsePlaintext{
		LoginCryptoResponse: &pb.LoginCryptoResponseUnion{
			DiffieHellman: &pb.LoginCryptoDiffieHellmanResponse{
				Hmac: sharedKeys.Challenge,
			},
		},
		PowResponse:    &pb.PoWResponseUnion{},
		CryptoResponse: &pb.CryptoResponseUnion{},
	}

	clientResRaw, err := proto.Marshal(clientRes)
	if err != nil {
		return nil, err
	}
	if _, err := c.send([]byte{}, clientResRaw); err != nil {
		return nil, err
	}

	return &sharedKeys, nil
}
func (p *insecureStream) receive() ([]byte, error) {
	var size uint32
	if err := binary.Read(p.stream, binary.BigEndian, &size); err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	binary.BigEndian.PutUint32(buf, size)
	if _, err := io.ReadFull(p.stream, buf[4:]); err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *insecureStream) send(prefix, data []byte) ([]byte, error) {
	var buf bytes.Buffer
	buf.Write(prefix)
	binary.Write(&buf, binary.BigEndian, uint32(len(prefix)+4+len(data)))
	buf.Write(data)

	pkt := buf.Bytes()

	_, err := c.stream.Write(pkt)
	return pkt, err
}

func makeHelloMessage(publicKey []byte, nonce []byte) []byte {
	hello := &pb.ClientHello{
		BuildInfo: &pb.BuildInfo{
			Product:  pb.Product_PRODUCT_PARTNER.Enum(),
			Platform: pb.Platform_PLATFORM_LINUX_X86.Enum(),
			Version:  proto.Uint64(0x10800000000),
		},
		CryptosuitesSupported: []pb.Cryptosuite{
			pb.Cryptosuite_CRYPTO_SUITE_SHANNON},
		LoginCryptoHello: &pb.LoginCryptoHelloUnion{
			DiffieHellman: &pb.LoginCryptoDiffieHellmanHello{
				Gc:              publicKey,
				ServerKeysKnown: proto.Uint32(1),
			},
		},
		ClientNonce: nonce,
		FeatureSet: &pb.FeatureSet{
			Autoupdate2: proto.Bool(true),
		},
		Padding: []byte{0x1e},
	}

	packetData, err := proto.Marshal(hello)
	if err != nil {
		log.Fatal("login marshaling error: ", err)
	}

	return packetData
}
