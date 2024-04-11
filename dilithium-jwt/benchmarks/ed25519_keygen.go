package main

import (
	"fmt"
	"runtime"
	"time"

	base "gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base"
	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

func ed25519KeyError(err error) {
	if err != nil {
		panic(err)
	}
}

// Generate a JWT claim
var ed25519TokenClaims = jwt.MapClaims{
	"username": "Hassan M.M",
	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	"iat":      time.Now().Unix(),
}

func main() {
	// Import ed25519 keys from storage
	ed25519PrivateKey, _ := base.ImportEd25519Keys()

	// Generate an ed25519 JWT token
	token, tokenErr := base.GenTokenEd25519(ed25519TokenClaims, ed25519PrivateKey)
	ed25519KeyError(tokenErr)

	// Get the current RAM
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory used for token generation (ED25519): %v KB\n", m.Alloc / 1024)

	// Export the Dilithium token
	base.ExportJWT(token, "ed25519")
}
