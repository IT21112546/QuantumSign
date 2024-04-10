package main

import (
	"fmt"
	"time"

	"runtime"

	base "gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base"
	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

func dilithiumKeyError(err error) {
	if err != nil {
		panic(err)
	}
}

// Generate a JWT claim
var dilithiumTokenClaims = jwt.MapClaims{
	"username": "Hassan M.M",
	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	"iat":      time.Now().Unix(),
}

func main() {
	// Import Dilithium keys from storage
	_, dilithiumPrivateKey, err := base.ImportDilithiumKeys()
	dilithiumKeyError(err)

	// Generate an dilithium JWT token
	_, tokenErr := base.GenTokenDilithium(dilithiumTokenClaims, dilithiumPrivateKey)
	dilithiumKeyError(tokenErr)

	// Get the current RAM
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory used for token generation (Dilithium): %v KB\n", m.Alloc / 1024)
}
