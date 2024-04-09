package algorithms

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

// A function to check for errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Generates an RSA key-pair
func GenRSA() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	checkErr(err)
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}

// A function to generate a JWT token with the given claims and private key
func GenTokenRSA(claims jwt.MapClaims, privateKey *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)

	// Benchmarking
	return signedToken, err
}

// A function to verify a JWT token using the given public key
func VerifyTokenRSA(signedToken string, publicKey *rsa.PublicKey) (bool, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func ExportRSAKeys(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) {
	// Export the private key
	privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		},
	)
	os.WriteFile("keys/rsa/id_rsa", privKeyPEM, 0600) // Make sure only the user can read and write

	// Export the public key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)
	os.WriteFile("keys/rsa/id_rsa.pub", publicKeyPEM, 0644)
}

func ImportRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	// Read the private key
	privKeyPEM, err := os.ReadFile("keys/rsa/id_rsa")
	checkErr(err)
	privKeyBlock, _ := pem.Decode(privKeyPEM)
	privKey, err := x509.ParsePKCS1PrivateKey(privKeyBlock.Bytes)
	checkErr(err)

	// Read the public key
	pubKeyPEM, err := os.ReadFile("keys/rsa/id_rsa.pub")
	checkErr(err)
	pubKeyBlock, _ := pem.Decode(pubKeyPEM)
	pubKey, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	checkErr(err)

	return privKey, pubKey.(*rsa.PublicKey)
}
