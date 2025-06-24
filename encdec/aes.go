package encdec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func AESEncrypt(plaintext string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("AESEncrypt.aes.NewCipher: %s", err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("AESEncrypt.cipher.NewGCM: %s", err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("AESEncrypt.io.ReadFull: %s", err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AESDecrypt(ciphertext string, key string) (string, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("AESDecrypt.base64.StdEncoding.DecodeString: %s", err.Error())
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("AESDecrypt.aes.NewCipher: %s", err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("AESDecrypt.cipher.NewGCM: %s", err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce := decodedCiphertext[:nonceSize]

	encryptedData := decodedCiphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", fmt.Errorf("AESDecrypt.aesGCM.Open: %s", err.Error())
	}

	return string(plaintext), nil
}
