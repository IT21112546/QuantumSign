package utils

import (
	"context"

	pb "github.com/qdrant/go-client/qdrant"
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
