package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// 0x6574685f345f3100000000000000000000000000000000000000000000000000
func StringToHex(raw []byte) [32]byte {
	result := fmt.Sprintf("0x%s", hex.EncodeToString(raw))
	lenPad := 32 - len(result)
	for lenPad > 0 {
		result += "0"
		lenPad -= 1
	}
	fmt.Println("len(nonce)", lenPad, result)

	nonce := [32]byte{}
	copy(nonce[:], result)

	return nonce
}

func PadBytes32(input [32]byte) string {
	builder := strings.Builder{}
	for _, runee := range input {
		builder.WriteByte(runee)
	}
	nonce := builder.String()

	padLen := 66 - 32
	for padLen > 0 {
		nonce += "0"
		padLen -= 1
	}

	return nonce
}

func main() {
	res := StringToHex([]byte("eth_4_1"))
	nonce := PadBytes32(res)

	fmt.Println(nonce, "0x6574685f345f3100000000000000000000000000000000000000000000000000 == nonce", "0x6574685f345f3100000000000000000000000000000000000000000000000000" == nonce)
}
