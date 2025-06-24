package main

import (
	"flag"
	"fmt"

	"github.com/flazhgrowth/fg-gotools/encdec"
	"github.com/flazhgrowth/fg-gotools/password"
)

func main() {
	opt := flag.Int("mode", 0, "")
	val := flag.String("val", "", "")
	plain := flag.String("plain", "", "")
	hashed := flag.String("hashed", "", "")
	salt := flag.String("salt", "", "")
	flag.Parse()

	if *opt == 0 {
		return
	}
	if *opt == 1 {
		encrypt(*val)
	} else if *opt == 2 {
		decrypt(*plain, *hashed, *salt)
	} else if *opt == 3 {
		encryptAES(*plain, *salt)
	} else if *opt == 4 {
		decryptAES(*hashed, *salt)
	}
}

func encrypt(val string) {
	hashed, salt, _ := password.Build(val)
	fmt.Printf("Hashed: %s (%s)\n", hashed, salt)
}

func decrypt(plain, hashed, salt string) {
	fmt.Println("Result", password.Assert(plain, hashed, salt))
}

func encryptAES(val string, key string) {
	cipher, err := encdec.AESEncrypt(val, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ciphertext", cipher)
}

func decryptAES(cipher, key string) {
	plain, err := encdec.AESDecrypt(cipher, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain", plain)
}
