package encoding

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Demo() {
	var message = []byte{8, 9, 10, 11, 12, 20, 255}
	fmt.Printf("message used %v\n", message)
	fmt.Println()
	var encoded = make([]byte, hex.EncodedLen(len(message)))
	hex.Encode(encoded, message)
	fmt.Printf("HEX of %b in decimal %v\n", message, encoded)
	fmt.Printf("HEX of %b in hex %s\n", message, encoded)
	fmt.Println()
	fmt.Printf("Base64 of %b is %s\n", message, base64.StdEncoding.EncodeToString(message))
	fmt.Printf("Base64 URL Safe of %b is %s\n", message, base64.URLEncoding.EncodeToString(message))
	fmt.Printf("Base64 URL Safe No Padding of %b is %s\n", message, base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(message))
	fmt.Println()
	fmt.Printf("Base32 of %b is %s\n", message, base32.StdEncoding.EncodeToString(message))
	fmt.Println()
}
