package crypto

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

type Box struct {
	aead cipher.AEAD
}

func NewBox(key []byte) (*Box, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key length: %d", len(key))
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	return &Box{aead: aead}, nil
}

func (b *Box) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := b.aead.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

func (b *Box) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < chacha20poly1305.NonceSizeX {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:chacha20poly1305.NonceSizeX]
	body := ciphertext[chacha20poly1305.NonceSizeX:]

	return b.aead.Open(nil, nonce, body, nil)
}
