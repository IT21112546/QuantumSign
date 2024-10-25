package algorithms

import (
	"fmt"
	"os"
	"time"

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

// Verifies a JWT token using the Dilithium public key and checks for expiration
func VerifyTokenDilithium(signedToken string, publicKey dilithium.PublicKey) (bool, bool, error) {
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token uses the Dilithium signing method
		if _, ok := token.Method.(*jwt.SigningMethodDilithium); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the public key for signature verification
		return publicKey, nil
	})

	if err != nil {
		return false, false, fmt.Errorf("error parsing token: %v", err)
	}

	// Check if the token is valid and contains claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			return time.Now().After(expirationTime), true, nil
		}
	}

	// If the token is invalid or claims are not present
	return false, false, fmt.Errorf("invalid token")
}

func ImportDilithiumKeys() (dilithium.PublicKey, dilithium.PrivateKey, error) {
	// Read the public key bytes from file
	publicKeyBytes, err := os.ReadFile("keys/id_dilithium.pub")
	if err != nil {
		return nil, nil, err
	}

	// Read the private key bytes from file
	privateKeyBytes, err := os.ReadFile("keys/id_dilithium")
	if err != nil {
		return nil, nil, err
	}

	// Create a new Dilithium public key from the imported bytes
	publicKey := dilithium.Mode3.PublicKeyFromBytes(publicKeyBytes)

	// Create a new Dilithium private key from the imported bytes
	privateKey := dilithium.Mode3.PrivateKeyFromBytes(privateKeyBytes)

	return publicKey, privateKey, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
