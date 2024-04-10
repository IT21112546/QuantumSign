package main

import (
	"fmt"

	"runtime"

	base "gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base"
)

func rsaError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	token, importErr := base.ImportJWT("rsa")
	_, rsaPublicKey := base.ImportRSAKeys()
	rsaError(importErr)

	_, verifyError := base.VerifyTokenRSA(token, rsaPublicKey)
	rsaError(verifyError)

	// Get the current RAM
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory used for verification (RS256): %v KB\n\n", m.Alloc / 1024)

}
