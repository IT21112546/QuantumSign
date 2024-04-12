package algorithms

import (
	"fmt"
	"os"

	"github.com/cloudflare/circl/sign/dilithium"
	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

// Generates a Dilithium key-pair
func GenDilithium() (dilithium.PublicKey, dilithium.PrivateKey) {
	publicKey, privateKey, err := dilithium.Mode3.GenerateKey(nil)
	checkErr(err)

	return publicKey, privateKey
}

// Generates a JWT token using the Dilithium private key
func GenTokenDilithium(claims jwt.MapClaims, privateKey dilithium.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodDilithium3, claims)
	signedToken, err := token.SignedString(privateKey)
	return signedToken, err
}

// Verifies a JWT token using the Dilithium public key
func VerifyTokenDilithium(signedToken string, publicKey dilithium.PublicKey) (bool, error) {
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodDilithium); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	checkErr(err)
	return parsedToken.Valid, nil
}

func ExportDilithiumKeys(privateKey dilithium.PrivateKey, publicKey dilithium.PublicKey) {
	// Export private key to file
	privateKeyFile, err := os.Create("keys/dilithium/id_dilithium")
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()
	_, err = privateKeyFile.Write(privateKey.Bytes())
	if err != nil {
		panic(err)
	}

	// Export public key to file
	publicKeyFile, err := os.Create("keys/dilithium/id_dilithium.pub")
	if err != nil {
		panic(err)
	}
	defer publicKeyFile.Close()
	_, err = publicKeyFile.Write(publicKey.Bytes())
	if err != nil {
		panic(err)
	}
}

func ImportDilithiumKeys() (dilithium.PublicKey, dilithium.PrivateKey, error) {
	// Read the public key bytes from file
	publicKeyBytes, err := os.ReadFile("keys/dilithium/id_dilithium.pub")
	if err != nil {
		return nil, nil, err
	}

	// Read the private key bytes from file
	privateKeyBytes, err := os.ReadFile("keys/dilithium/id_dilithium")
	if err != nil {
		return nil, nil, err
	}

	// Create a new Dilithium public key from the imported bytes
	publicKey := dilithium.Mode3.PublicKeyFromBytes(publicKeyBytes)

	// Create a new Dilithium private key from the imported bytes
	privateKey := dilithium.Mode3.PrivateKeyFromBytes(privateKeyBytes)

	return publicKey, privateKey, nil
}
