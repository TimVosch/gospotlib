package pb

import (
	"bytes"
	"crypto/aes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
	"google.golang.org/protobuf/proto"
)

func blobKey(username string, secret []byte) []byte {
	data := pbkdf2.Key(secret, []byte(username), 0x100, 20, sha1.New)[0:20]

	hash := sha1.Sum(data)
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, 20)
	return append(hash[:], length...)
}

func CredentialsFromBlob(blob64 []byte, username, deviceID string) (*LoginCredentials, error) {
	blobEnc, err := base64.StdEncoding.DecodeString(string(blob64))
	if err != nil {
		return nil, err
	}

	// Prepare decryption info
	secret := sha1.Sum([]byte(deviceID))
	key := blobKey(username, secret[:])

	// Decrypt
	blk, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create aes cipher for encrypted blob: %s", err)
	}

	blkSize := blk.BlockSize()
	if len(blobEnc)%blkSize != 0 {
		return nil, errors.New("invalid block size")
	}
	blob := make([]byte, len(blobEnc))
	for i := 0; i < len(blobEnc); i += blkSize {
		blk.Decrypt(blob[i:], blobEnc[i:])
	}

	l := len(blob)
	for i := 0; i < l-0x10; i++ {
		blob[l-i-1] ^= blob[l-i-0x11]
	}

	// Extract information from blob
	buf := bytes.NewBuffer(blob)
	// skip
	if _, err := buf.ReadByte(); err != nil {
		return nil, fmt.Errorf("failed to read byte from blob: %w", err)
	}
	// skip
	if _, err := readPrefixedBlock(buf); err != nil {
		return nil, fmt.Errorf("failed to read prefixed byte block from blob: %w", err)
	}
	// skip
	if _, err := buf.ReadByte(); err != nil {
		return nil, fmt.Errorf("failed to read byte from blob: %w", err)
	}
	// get authType
	authType, err := readShort(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read short from blob: %w", err)
	}
	// skip
	if _, err := buf.ReadByte(); err != nil {
		return nil, fmt.Errorf("failed to read byte from blob: %w", err)
	}
	// read authData
	authData, err := readPrefixedBlock(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read prefixed byte block from blob: %w", err)
	}

	return &LoginCredentials{
		Username: proto.String(username),
		Typ:      AuthenticationType(authType).Enum(),
		AuthData: authData,
	}, nil
}

func readShort(buf *bytes.Buffer) (uint16, error) {
	lo, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	if lo&0x80 == 0 {
		return uint16(lo), nil
	}

	hi, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}

	return (uint16(lo & 0x7F)) | uint16(hi<<7), nil
}

func readPrefixedBlock(buf *bytes.Buffer) ([]byte, error) {
	len, err := readShort(buf)
	if err != nil {
		return nil, err
	}

	data := make([]byte, len)
	if _, err := buf.Read(data); err != nil {
		return nil, err
	}

	return data, nil
}
