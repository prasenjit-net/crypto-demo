package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func Demo() {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	public := key.PublicKey
	message := "Some data to be enciphered"

	encrypted, _ := rsa.EncryptPKCS1v15(rand.Reader, &public, []byte(message))
	fmt.Printf("RSA encrypted data %v\n", encrypted)

	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, key, encrypted)
	if err != nil {
		fmt.Printf("failed to decrypt %v\n", err)
	} else {
		fmt.Printf("RSA decrypted value \"%s\"", decrypted)
	}

	// AES or symmetric encryption

	aesKey := make([]byte, 32)
	_, _ = rand.Reader.Read(aesKey)

	c, _ := aes.NewCipher(aesKey)
	gcm, _ := cipher.NewGCM(c)

	nonce := make([]byte, gcm.NonceSize())
	_, _ = rand.Reader.Read(nonce)

	encrypted = make([]byte, 0)
	decrypted = make([]byte, 0)

	fmt.Println()
	fmt.Println()

	encrypted = gcm.Seal(encrypted, nonce, []byte(message), nil)
	fmt.Printf("AES encrypted data %v\n", encrypted)

	decrypted, err = gcm.Open(decrypted, nonce, encrypted, nil)
	if err != nil {
		fmt.Printf("decrypt failed with error %v\n", err)
	} else {
		fmt.Printf("AES decrypted value is \"%s\"\n", decrypted)
	}
}
