package utils

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	pb "github.com/qdrant/go-client/qdrant"
)

// Create a new user
func CreateUserPoints(req CreateUserBody) ([]*pb.PointStruct, error) {

	// Check if user already exists
	res, isRegistered, _ := GetUserData(req.PubKey)
	if isRegistered {
		log.Printf("User with public key already exists: %s", res.Username)
		return nil, nil
	}

	id := uuid.NewString()
	upsertData := []*pb.PointStruct{
		{
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Uuid{
					Uuid: id,
				},
			},
			Vectors: &pb.Vectors{
				VectorsOptions: &pb.Vectors_Vector{
					Vector: &pb.Vector{
						Data: req.PubKey,
					},
				},
			},
			Payload: map[string]*pb.Value{
				"username": {
					Kind: &pb.Value_StringValue{StringValue: req.Username},
				},
				"email": {
					Kind: &pb.Value_StringValue{StringValue: req.Email},
				},
			},
		},
	}

	return upsertData, nil
}

// Upsert the points to Qdrant
func UpsertPoints(user []*pb.PointStruct) (error, float64) {
	conn := GetPointsClient()

	waitUpsert := true
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Save the points to the database
	res, err := conn.Upsert(ctx, &pb.UpsertPoints{
		CollectionName: COLLECTION_NAME,
		Wait:           &waitUpsert,
		Points:         user,
	})

	return err, res.Time * 1000
}

// Get all user data for a public key
func GetUserData(key []float32) (UserSearchResult, bool, float64) {
	conn := GetPointsClient()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	res, err := conn.Search(ctx, &pb.SearchPoints{
		CollectionName: COLLECTION_NAME,
		Vector:         key,
		Limit:          1,
		WithPayload:    &pb.WithPayloadSelector{SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true}},
	})
	CheckErr(err)

	// Preparing the results
	parsedSearchResult := parseSearchRes(res)
	isUserRegistered := validateScore(res)
	timeToSearch := res.Time * 1000

	return parsedSearchResult, isUserRegistered, timeToSearch
}

// Parse a search response and return a UserSearchResult
func parseSearchRes(res *pb.SearchResponse) UserSearchResult {
	searchStruct := UserSearchResult{}

	for _, point := range res.Result {
		for key, value := range point.Payload {
			switch key {
			case "username":
				searchStruct.Username = value.String()
			case "email":
				searchStruct.Email = value.String()
			}
		}
	}

	return searchStruct
}

// Check if the score is a perfect 0
func validateScore(res *pb.SearchResponse) bool {
	var score float32
	for _, point := range res.Result {
		score = point.Score
	}
	return score > THRESHOLD_SCORE
}
