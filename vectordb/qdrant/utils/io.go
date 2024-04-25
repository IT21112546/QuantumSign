package utils

import (
	"os"
	"unsafe"

	"github.com/cloudflare/circl/kem"
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

// Convert a float32 slice to a byte slice
func FloatToByte(in []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&in[0])), len(in)*4)
}

// Export a public key to file
func ExportPubKey(pub kem.PublicKey, filename string) {
	binaryPub, err := pub.MarshalBinary()
	CheckErr(err)
	os.WriteFile(filename, binaryPub, 0644)
}

// Import a pubilc key to a float32 array
func ImportPubKey(filename string) []float32 {
	bin, err := os.ReadFile(filename)
	CheckErr(err)

	pub, err := kyber512.Scheme().UnmarshalBinaryPublicKey(bin)
	CheckErr(err)

	pubBin, err := pub.MarshalBinary()
	CheckErr(err)

	return ByteToFloat(pubBin)
}
