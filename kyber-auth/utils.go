package utils

import (
	"log"
	"os"

	"github.com/cloudflare/circl/kem"
	"github.com/cloudflare/circl/kem/kyber/kyber512"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func CyclicXOR(a, b []byte) []byte {
	// Extend B cyclically to match the length of A
	extendedB := make([]byte, len(a))
	for i := range extendedB {
		extendedB[i] = b[i%len(b)]
	}

	// XOR each byte of A with the corresponding byte from the extended B
	encryptedA := make([]byte, len(a))
	for i := range a {
		encryptedA[i] = a[i] ^ extendedB[i]
	}

	return encryptedA
}

// Import a private key
func ImportPriKey(filename string) (kem.PrivateKey) {
	rawPrivateKey, err := os.ReadFile(filename)
	privateKey, err := kyber512.Scheme().UnmarshalBinaryPrivateKey(rawPrivateKey)
	CheckErr(err)
	return privateKey
}

func ImportPubKey(filename string) (kem.PublicKey) {
	rawPublicKey, err := os.ReadFile(filename)
	publicKey, err := kyber512.Scheme().UnmarshalBinaryPublicKey(rawPublicKey)
	CheckErr(err)
	return publicKey
}
