package password

import (
	"crypto/rand"
	"encoding/base64"
	"unicode"

	"github.com/flazhgrowth/fg-gotools/hash/argon"
	"github.com/flazhgrowth/fg-gotools/hash/bcrypt"
)

func Build(password string) (hashedPassword, salt string, err error) {
	bsalt := make([]byte, saltLength)
	if _, err = rand.Read(bsalt); err != nil {
		return
	}

	hashedPassword = argon.Hash(password, bsalt)
	hashedPassword, err = bcrypt.Hash(hashedPassword)
	salt = base64.RawStdEncoding.EncodeToString(bsalt)

	return
}

func Assert(plainPassword, hashedPassword, salt string) bool {
	bsalt, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return false
	}
	return bcrypt.Assert(argon.Hash(plainPassword, bsalt), hashedPassword)
}

func IsPasswordStrength(password string) bool {
	if len(password) < 7 {
		return false
	}

	var hasNumber, hasUpper, hasLower, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasNumber && hasUpper && hasLower && hasSpecial
}
