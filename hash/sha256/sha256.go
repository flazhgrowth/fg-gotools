package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
