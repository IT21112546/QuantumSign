package utils

import (
	"context"
	"fmt"

	pb "github.com/qdrant/go-client/qdrant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// A function to create a new gRCP connection
func CreateGrpcConnection() *grpc.ClientConn {
	conn, err := grpc.DialContext(
		context.Background(),
		ADDR,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	CheckErr(err)
	return conn
}

// A function to create a new Qdrant points client
func CreatePointsClient(conn *grpc.ClientConn) pb.PointsClient {
	return pb.NewPointsClient(conn)
}

// Establish global connections
var GrpcConn = CreateGrpcConnection()

// Getters for the global gRPC connection
func GetGrpcConn() *grpc.ClientConn {
	return GrpcConn
}

// Create a points client
func GetPointsClient() pb.PointsClient {
	return CreatePointsClient(GrpcConn)
}

// Create a collection client
func GetCollectionsClient() pb.CollectionsClient {
	return pb.NewCollectionsClient(GrpcConn)
}

// MongoDB client and collection (global variables)
var mongoClient *mongo.Client
var clientsCollection *mongo.Collection

// Connect to MongoDB
func InitMongoDB() {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_ADDR))
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
	} else {
		fmt.Println("Connected to MongoDB:", MONGO_ADDR)
	}
	clientsCollection = mongoClient.Database(MONGO_DATABASE).Collection(MONGO_COLLECTION)
}
