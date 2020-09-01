package digest

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func Demo() {
	sum1 := sha1.Sum([]byte("some interesting message"))
	sum2 := sha1.Sum([]byte("some interesting message"))

	hex1 := hex.EncodeToString(sum1[:])
	hex2 := hex.EncodeToString(sum2[:])

	fmt.Printf(" first sum is: %s\n", hex1)
	fmt.Printf("second sum is: %s\n", hex2)

	fmt.Printf("length of SHA1 is %d\n", sha1.Size)
	fmt.Printf("length of SHA256 is %d\n", sha256.Size)
	fmt.Printf("length of SHA512 is %d\n", sha512.Size)
	fmt.Printf("length of MD5 is %d\n", md5.Size)
}
