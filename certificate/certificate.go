package certificate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"github.com/grantae/certinfo"
	"math/big"
	rand2 "math/rand"
	"net"
	"time"
)

func Demo() {
	certifier := &x509.Certificate{
		SerialNumber: big.NewInt(rand2.Int63()),
		Subject: pkix.Name{
			CommonName:         "Authority",
			OrganizationalUnit: []string{"it", "india"},
		},
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365 * 2),
		KeyUsage:              x509.KeyUsageCertSign,
		IsCA:                  true,
	}

	leaf := &x509.Certificate{
		SerialNumber: big.NewInt(rand2.Int63()),
		Subject: pkix.Name{
			CommonName:         "localhost",
			OrganizationalUnit: []string{"hr", "india"},
		},
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 180),
		KeyUsage:              x509.KeyUsageDataEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:              []string{"localhost"},
		IPAddresses: []net.IP{
			net.ParseIP("127.0.0.1"),
		},
	}
	caKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	leafKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	publicKey := caKey.PublicKey
	caBytes, _ := x509.CreateCertificate(rand.Reader, certifier, certifier, &publicKey, caKey)
	caCert, _ := x509.ParseCertificate(caBytes)

	publicKey = leafKey.PublicKey
	leafBytes, _ := x509.CreateCertificate(rand.Reader, leaf, caCert, &publicKey, caKey)
	leafCert, _ := x509.ParseCertificate(leafBytes)

	text, _ := certinfo.CertificateText(caCert)
	fmt.Printf("==================CA START==================\n")
	println(text)
	fmt.Printf("===================CA END===================\n")
	text, _ = certinfo.CertificateText(leafCert)
	fmt.Printf("=================leaf START=================\n")
	println(text)
	fmt.Printf("==================leaf END==================\n")

	roots := x509.NewCertPool()
	roots.AddCert(caCert)

	opts := x509.VerifyOptions{
		DNSName: "localhost",
		Roots:   roots,
	}
	chain, err := leafCert.Verify(opts)
	if err != nil {
		fmt.Printf("Invalid cert %v", err)
	} else {
		fmt.Printf("%d certificates in chain", len(chain))
	}
}
