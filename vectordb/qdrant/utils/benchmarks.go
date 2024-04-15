package utils

import (
	"log"

	"github.com/bxcodec/faker/v3"
)

// Populate an array of public keys
func GenPubKeys(count int) [][]float32 {
	pubKeys := make([][]float32, count)
	for i := range pubKeys {
		pubKey, _ := GenKeyPair()
		pubByte, _ := pubKey.MarshalBinary()
		pubVector := ByteToFloat(pubByte)
		pubKeys[i] = pubVector
	}
	return pubKeys
}

func BenchmarkAverages(pubKeys [][]float32, batchSize int) {
	var writeAvg float64

	for i, v := range pubKeys {
		// Create a user
		pointsStruct, _ := CreateUserPoints(CreateUserBody{
			PubKey:   v,
			Username: faker.Username(),
			Email:    faker.Email(),
		})

		_, insertTime := UpsertPoints(pointsStruct)
		writeAvg += insertTime

		// For every sampleSize iterations
		if i >= batchSize && i % batchSize == 0 {
			log.Println(i)
			writeAvg = writeAvg / float64(BENCHMARK_ITERATIONS)
			log.Printf("Average write time per %d iterations: %.3f ms", i, writeAvg)

			// Perform a read
			_, _, searchTime := GetUserData(v)
			log.Printf("Average read time per %d iterations: %.3f ms", i, searchTime)
		}
	}
}
