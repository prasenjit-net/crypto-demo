package password

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Demo() {
	var password = "may very hard password"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 13)
	fmt.Printf("Hashed password to store %s\n", hashed)

	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err == nil {
		fmt.Printf("Password match with %s\n", password)
	}
	wrong := []byte("password")
	err = bcrypt.CompareHashAndPassword(hashed, wrong)
	if err != nil {
		fmt.Printf("Wrong password doesnt match with %s\n", wrong)
	}

	start := time.Now()
	var i int
	for i = 0; i < 10; i++ {
		_ = bcrypt.CompareHashAndPassword(hashed, []byte(password))
	}
	fmt.Printf("BCRYPT took %v for 10 times\n", time.Now().Sub(start))

	start = time.Now()
	for i = 0; i < 1000; i++ {
		_ = sha256.Sum256(hashed)
	}
	fmt.Printf("SHA256 took %v for 1000 times\n", time.Now().Sub(start))
}
