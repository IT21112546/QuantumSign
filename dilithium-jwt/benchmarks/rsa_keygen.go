package main

import (
	"fmt"
	"time"

	"runtime"

	base "gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base"
	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

func rsaKeyError(err error) {
	if err != nil {
		panic(err)
	}
}

// Generate a JWT claim
var rsaTokenClaims = jwt.MapClaims{
	"username": "Hassan M.M",
	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	"iat":      time.Now().Unix(),
}

func main() {
	// Import RSA keys from storage
	rsaPrivateKey, _ := base.ImportRSAKeys()

	// Generate an RSA JWT token
	_, err := base.GenTokenRSA(rsaTokenClaims, rsaPrivateKey)
	rsaKeyError(err)

	// Get the current RAM
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory used for token generation (RSA): %v KB\n", m.Alloc / 1024)
}
