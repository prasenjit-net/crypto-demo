package main

import (
	"flag"
	"fmt"
	"github.com/prasenjit-net/crypto-demo/certificate"
	"github.com/prasenjit-net/crypto-demo/digest"
	"github.com/prasenjit-net/crypto-demo/encoding"
	"github.com/prasenjit-net/crypto-demo/encryption"
	"github.com/prasenjit-net/crypto-demo/mac"
	"github.com/prasenjit-net/crypto-demo/password"
	"github.com/prasenjit-net/crypto-demo/signature"
)

var demoFunction string

func main() {
	flag.StringVar(&demoFunction, "demo", "digest", "give the demo name to run")
	flag.Parse()

	switch demoFunction {
	case "digest":
		digest.Demo()
	case "encryption":
		encryption.Demo()
	case "signature":
		signature.Demo()
	case "mac":
		mac.Demo()
	case "encoding":
		encoding.Demo()
	case "password":
		password.Demo()
	case "certificate":
		certificate.Demo()
	default:
		fmt.Println("Invalid demo name")
	}
}
