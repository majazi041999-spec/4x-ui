package quic

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidECHKeyLength = errors.New("invalid ECH key length")
	ErrCiphertextTooShort  = errors.New("ciphertext too short")
)

// ECHSuite provides minimal SNI sealing/opening primitives used by Task 1.2
// until wire-level ECH integration is added.
type ECHSuite struct {
	key []byte
}

// GenerateECHKey creates a random 32-byte key (AES-256-GCM).
func GenerateECHKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("generate ECH key: %w", err)
	}
	return key, nil
}

// NewECHSuite builds a suite from a raw AES-256 key.
func NewECHSuite(key []byte) (*ECHSuite, error) {
	if len(key) != 32 {
		return nil, ErrInvalidECHKeyLength
	}
	keyCopy := make([]byte, len(key))
	copy(keyCopy, key)
	return &ECHSuite{key: keyCopy}, nil
}

// SealSNI encrypts the given server name and returns base64 ciphertext.
func (e *ECHSuite) SealSNI(sni string) (string, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", fmt.Errorf("create cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("create gcm: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("generate nonce: %w", err)
	}

	sealed := gcm.Seal(nil, nonce, []byte(sni), nil)
	out := append(nonce, sealed...)
	return base64.RawStdEncoding.EncodeToString(out), nil
}

// OpenSNI decrypts a base64 ciphertext produced by SealSNI.
func (e *ECHSuite) OpenSNI(ciphertext string) (string, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", fmt.Errorf("create cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("create gcm: %w", err)
	}

	raw, err := base64.RawStdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("decode ciphertext: %w", err)
	}
	if len(raw) < gcm.NonceSize() {
		return "", ErrCiphertextTooShort
	}

	nonce := raw[:gcm.NonceSize()]
	payload := raw[gcm.NonceSize():]
	plain, err := gcm.Open(nil, nonce, payload, nil)
	if err != nil {
		return "", fmt.Errorf("decrypt sni: %w", err)
	}
	return string(plain), nil
}
