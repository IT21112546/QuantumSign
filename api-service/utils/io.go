package utils

import (
	"encoding/base64"
	"errors"

	"github.com/cloudflare/circl/kem/kyber/kyber512"
)

// Convert a byte slice to a float32 slice
func ByteToFloat(src []byte) []float32 {
	floatArr := make([]float32, len(src))

	for k, v := range src {
		floatArr[k] = float32(v)
	}

	return floatArr
}

// Import a pubilc key to a float32 array
func LoadPubKey(base64EncodedPubKey string) ([]float32, error) {
	bin, err := base64.StdEncoding.DecodeString(base64EncodedPubKey)

	// If the base64 decoding fails, return an error
	if err != nil {
		return nil, errors.New(BASE64_DECODING_ERROR_MSG)
	}

	pub, err := kyber512.Scheme().UnmarshalBinaryPublicKey(bin)

	// If the unmarshalling fails, return an error
	if err != nil {
		return nil, errors.New(ERROR_UNMARSHALLING_MSG)
	}

	// Convert the public key to a byte slice and then to a float32 slice
	pubBin, err := pub.MarshalBinary()

	return ByteToFloat(pubBin), nil
}
