package algorithms

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

// Generates an Ed25519 key-pair
func GenEd25519() (ed25519.PrivateKey, ed25519.PublicKey) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Error generating Ed25519 key pair: %v\n", err)
		return nil, nil
	}
	return privateKey, publicKey
}

// Generates a JWT token using the Ed25519 private key
func GenTokenEd25519(claims jwt.MapClaims, privateKey ed25519.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedToken, err := token.SignedString(privateKey)

	return signedToken, err
}

// Verifies a JWT token using the Ed25519 public key
func VerifyTokenEd25519(tokenString string, publicKey ed25519.PublicKey) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func ExportEd25519Keys(privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) {
	// Export the private key
	privKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	privKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: privKeyBytes,
		},
	)
	os.WriteFile("keys/ed25519/id_ed25519", privKeyPEM, 0600) // Ensure file is only accessible by the user

	// Export the public key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)
	os.WriteFile("keys/ed25519/id_ed25519.pub", publicKeyPEM, 0644)
}

func ImportEd25519Keys() (ed25519.PrivateKey, ed25519.PublicKey) {
	// Read the private key
	privKeyPEM, err := os.ReadFile("keys/ed25519/id_ed25519")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(privKeyPEM)
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// Read the public key
	pubKeyPEM, err := os.ReadFile("keys/ed25519/id_ed25519.pub")
	if err != nil {
		panic(err)
	}
	block, _ = pem.Decode(pubKeyPEM)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return privateKey.(ed25519.PrivateKey), publicKey.(ed25519.PublicKey)
}
