package utils

import (
	"context"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
)

// Create a new collection
func NewQdrantCollection() (err error) {

	// Create gRPC collection client
	client := GetCollectionsClient()

	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var defaultSegmentNumber uint64 = 2
	_, err = client.Create(ctx, &pb.CreateCollection{
		CollectionName: COLLECTION_NAME,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     uint64(VECTOR_SIZE),
				Distance: DISTANCE,
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &defaultSegmentNumber,
		},
	})

	return err
}

// Create a new keyword index
func NewKeywordIndex() (err error) {

	// Create a points client
	conn := GetPointsClient()

	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Create an index for each key
	for _, index := range Indexes {

		// Specify the field index
		fieldIndex1Type := pb.FieldType_FieldTypeKeyword
		fieldIndex1Name := index

		// Create the field index
		_, err = conn.CreateFieldIndex(ctx, &pb.CreateFieldIndexCollection{
			CollectionName: COLLECTION_NAME,
			FieldName:      fieldIndex1Name,
			FieldType:      &fieldIndex1Type,
		})
	}

	return err
}
