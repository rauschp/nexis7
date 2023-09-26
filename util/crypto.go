package util

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const PrivateKeyPath string = "keys/nexis.key"

type PrivateKey struct {
	Key ed25519.PrivateKey
}

func GenerateNewPrivateKey() *PrivateKey {
	_, priv, _ := ed25519.GenerateKey(rand.Reader)

	return &PrivateKey{
		Key: priv,
	}
}

func GeneratePrivateKeyFromBase64(key string) (*PrivateKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	ecKey := ed25519.PrivateKey(keyBytes)

	return &PrivateKey{
		Key: ecKey,
	}, nil
}

func (p *PrivateKey) Bytes() []byte {
	return p.Key
}

func (p *PrivateKey) ToBase64() string {
	return base64.StdEncoding.EncodeToString(p.Key)
}

func (p *PrivateKey) Sign(message []byte) []byte {
	return ed25519.Sign(p.Key, message)
}

// Public Key
type PublicKey struct {
	Key ed25519.PublicKey
}

type Address struct {
	Value []byte
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, 32)
	copy(b, p.Key[32:])

	return &PublicKey{
		Key: b,
	}
}

func (p *PublicKey) GetAddress() Address {
	return Address{
		Value: p.Key[len(p.Key)-20:],
	}
}

func (a Address) ToBytes() []byte {
	return a.Value
}

func (a Address) ToString() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(a.Value))
}

func (p *PublicKey) Verify(message []byte, signature []byte) bool {
	return ed25519.Verify(p.Key, message, signature)
}
