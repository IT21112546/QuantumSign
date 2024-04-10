package main

import (
	"fmt"
	"runtime"

	base "gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base"
)

func ed25519Error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	token, importErr := base.ImportJWT("ed25519")
	_, ed25519PublicKey := base.ImportEd25519Keys()
	ed25519Error(importErr)

	_, verifyError := base.VerifyTokenEd25519(token, ed25519PublicKey)
	ed25519Error(verifyError)

	// Get the current RAM
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory used for verification (ED25519): %v KB\n", m.Alloc / 1024)
}
