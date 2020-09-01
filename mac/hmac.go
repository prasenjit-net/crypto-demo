package mac

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func Demo() {
	key := "secret hmac key"
	message := "message data"
	sum := computeHmac(key, message)
	fmt.Printf("key: %s data: %s generated mac is %v\n", key, message, sum)
	sum = computeHmac(key, message)
	fmt.Printf("key: %s data: %s generated mac is %v\n", key, message, sum)
	sum = computeHmac("key", message)
	fmt.Printf("key: %s data: %s generated mac is %v\n", key, message, sum)
}

func computeHmac(key string, message string) []byte {
	hash := hmac.New(sha1.New, []byte(key))
	sum := hash.Sum([]byte(message))
	return sum
}
