package main

import (
	"fmt"
	"runtime"

	base "gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base"
)

func dilithiumError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	token, importErr := base.ImportJWT("dilithium")
	dilithiumPublicKey, _, keyErr := base.ImportDilithiumKeys()
	dilithiumError(importErr)
	dilithiumError(keyErr)

	_, verifyError := base.VerifyTokenDilithium(token, dilithiumPublicKey)
	dilithiumError(verifyError)

	// Get the current RAM
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory used for verification (Dilithium): %v KB\n", m.Alloc / 1024)
}
