package main

import (
	"fmt"
	"time"

	base "gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base"
	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// A global variable to store the JWT token / Errors
var globalToken string
var err error

// Generate a JWT claim
var claims = jwt.MapClaims{
	"username": "Hassan M.M",
	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	"iat":      time.Now().Unix(),
}

// The function used to benchmark
func iterateBenchmark(iterations int, msg string, f func()) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		f()
	}
	elapsed := time.Since(start)
	average := elapsed / time.Duration(iterations)
	fmt.Printf("%s: %v\n", msg, average)
}

func main() {
	// Import Dilithium keys from storage
	dilPublicKey, dilPrivateKey, err := base.ImportDilithiumKeys()
	checkErr(err)

	// Generate Dilithium JWT token
	iterateBenchmark(1000, "Time to generate a Dilithium JWT token", func() {
		globalToken, err = base.GenTokenDilithium(claims, dilPrivateKey)
	})

	// Benchmark: Verify Dilithium JWT token
	iterateBenchmark(1000, "Time to verify a Dilithium JWT token", func() {
		_, err = base.VerifyTokenDilithium(globalToken, dilPublicKey)
	})
	fmt.Println()

	// Import RSA keys from storage
	rsaPrivateKey, rsaPublicKey := base.ImportRSAKeys()

	// Benchmark: Generate an RSA JWT token
	iterateBenchmark(1000, "Time to generate an RSA JWT token", func() {
		token, err := base.GenTokenRSA(claims, rsaPrivateKey)
		checkErr(err)
		globalToken = token
	})

	// Benchmark: Verify an RSA JWT token
	iterateBenchmark(1000, "Time to verify an RSA JWT token", func() {
		_, err = base.VerifyTokenRSA(globalToken, rsaPublicKey)
		checkErr(err)
	})
	fmt.Println()

	// Benchmark: Generate & verify Ed25519 JWT token
	ed25519PrivateKey, ed25519publicKey := base.ImportEd25519Keys()
	iterateBenchmark(1000, "Time to generate a Ed25519 JWT token", func() {
		globalToken, err = base.GenTokenEd25519(claims, ed25519PrivateKey)
	})

	// Benchmark: Verify Ed25519 JWT token
	iterateBenchmark(1000, "Time to verify an Ed25519 JWT token", func() {
		_, err = base.VerifyTokenEd25519(globalToken, ed25519publicKey)
	})
}
