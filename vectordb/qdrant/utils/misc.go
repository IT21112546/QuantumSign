package utils

import (
	"github.com/cloudflare/circl/kem"
	"github.com/cloudflare/circl/kem/kyber/kyber512"
)

// A helper function to create a key pair
func GenKeyPair() (kem.PublicKey, kem.PrivateKey) {
	kyber := kyber512.Scheme()
	pub, pri, err := kyber.GenerateKeyPair()
	CheckErr(err)
	return pub, pri
}

// To check for errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
