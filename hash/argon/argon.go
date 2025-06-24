package argon

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

const (
	Time    = 1
	Memory  = 64 * 1024
	KeyLen  = 32
	Threads = 4
)

func Hash(plain string, salt []byte) string {
	hash := argon2.IDKey([]byte(plain), salt, Time, Memory, Threads, KeyLen)
	return base64.RawStdEncoding.EncodeToString(hash)
}
