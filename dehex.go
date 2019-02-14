package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
)

func hexTob64Str(hexStr string) string {
	src := []byte(hexStr)
	dst := make([]byte, hex.DecodedLen(len(src)))

	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(dst[:n])
}

func b64ToDecimal(base64Str string) string {
	bytes, _ := base64.RawURLEncoding.DecodeString(removeEquals(base64Str))
	bigInt := big.Int{}
	bigInt.SetBytes(bytes)

	return bigInt.String()
}

func removeEquals(input string) string {
	//remove equals from string
	return input[:len(input)-1]
}

func main() {
	args := os.Args[1:]
	str := args[0]
	fmt.Printf("hex input: %s\n", str)
	b64Str := hexTob64Str(str)
	decimal := b64ToDecimal(b64Str)
	b64NoEquals := removeEquals(b64Str)

	fmt.Println(fmt.Sprintf(`"attributes": { 
		"b64": "%s", 
		"b64_std": "%s", 
		"b64_url": "%s", 
		"byte_length": "2", 
		"dec": "%s", 
		"hex": "%s", 
		"id": "%s"
	},`, b64NoEquals, b64Str, b64NoEquals, decimal, str, b64NoEquals))

}
