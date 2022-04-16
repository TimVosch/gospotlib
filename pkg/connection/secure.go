package connection

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/timvosch/gospot/internal/crypto"
)

// TODO: Implement locking for concurrent write on socket

const (
	SHN_MAC_LEN = 4
)

var (
	ErrInvalidMac = errors.New("invalid mac")
)

// secureStream ...
type secureStream struct {
	stream io.ReadWriteCloser

	sendNonce int
	recvNone  int

	sendCipher shn_ctx
	recvCipher shn_ctx
}

func newSecureStream(stream io.ReadWriteCloser, keys *crypto.SharedKeys) *secureStream {
	s := &secureStream{
		stream: stream,
	}

	setKey(&s.sendCipher, keys.SendKey)
	setKey(&s.recvCipher, keys.RecvKey)

	return s
}

func (s *secureStream) Send(cmd uint8, data []byte) error {
	// Build packet
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, cmd)
	binary.Write(&buf, binary.BigEndian, uint16(len(data)))
	buf.Write(data)

	// Encrypt
	encrypted := buf.Bytes()
	shn_encrypt(&s.sendCipher, encrypted, len(encrypted))

	// Calculate mac
	mac := make([]byte, SHN_MAC_LEN)
	shn_finish(&s.sendCipher, mac, len(mac))
	s.sendNonce++
	setNonce(&s.sendCipher, s.sendNonce)

	// Send
	_, err := s.stream.Write(append(encrypted, mac...))

	return err
}

func (s *secureStream) Receive() (uint8, []byte, error) {
	// Read header
	header := make([]byte, 3)
	if _, err := io.ReadFull(s.stream, header); err != nil {
		return 0, nil, err
	}
	// Decrypt header
	shn_decrypt(&s.recvCipher, header, len(header))

	cmd := header[0]
	size := binary.BigEndian.Uint16(header[1:])

	data := make([]byte, size+SHN_MAC_LEN)
	if _, err := io.ReadFull(s.stream, data); err != nil {
		return 0, nil, fmt.Errorf("error reading data: %w", err)
	}

	// Decrypt
	decrypted := data[:size]
	shn_decrypt(&s.recvCipher, decrypted, int(size))
	mac := data[size:]
	cmpMac := make([]byte, SHN_MAC_LEN)
	shn_finish(&s.recvCipher, cmpMac, len(cmpMac))

	// Check mac
	if !bytes.Equal(mac, cmpMac) {
		return 0, nil, ErrInvalidMac
	}

	// Update nonce
	s.recvNone++
	setNonce(&s.recvCipher, s.recvNone)

	return cmd, decrypted, nil
}

func (s *secureStream) Close() {
	s.stream.Close()
}

func setKey(ctx *shn_ctx, key []uint8) {
	shn_key(ctx, key, len(key))
	setNonce(ctx, 0)
}

func setNonce(ctx *shn_ctx, nonce int) {
	n := make([]byte, 4)
	binary.BigEndian.PutUint32(n, uint32(nonce))
	shn_nonce(ctx, n, len(n))
}
