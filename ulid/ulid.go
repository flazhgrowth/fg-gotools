package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func Generate() (string, error) {
	// Set a secure random source
	entropy := ulid.Monotonic(rand.Reader, 0)

	// Get current timestamp
	t := time.Now().UTC()

	// Generate ULID
	id, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
