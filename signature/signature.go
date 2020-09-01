package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func Demo() {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	public := key.PublicKey
	message := "Some data to be signed"

	hashOfData := sha256.Sum256([]byte(message))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashOfData[:])
	fmt.Printf("Signature of data %X\n", signature)

	err := rsa.VerifyPKCS1v15(&public, crypto.SHA256, hashOfData[:], signature)
	if err != nil {
		fmt.Printf("signature is invalid with error %v\n", err)
	} else {
		fmt.Println("signature is valid")
	}
}
