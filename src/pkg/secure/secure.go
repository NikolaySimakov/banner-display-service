package secure

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type APISecure interface {
	Hash(key string) string
	GenerateKey() string
}

func NewSecure(salt string) *Secure {
	return &Secure{salt: salt}
}

type Secure struct {
	salt string
}

func (s *Secure) Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}

func (s *Secure) GenerateKey() string {
	keyBytes := make([]byte, 32)
	_, _ = rand.Read(keyBytes)
	return hex.EncodeToString(keyBytes)
}