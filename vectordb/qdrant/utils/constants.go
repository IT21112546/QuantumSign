package utils

import (
	pb "github.com/qdrant/go-client/qdrant"
)

// Qdrant database parameters
const ADDR = "localhost:6334"
const COLLECTION_NAME = "pqc"
const VECTOR_SIZE uint32 = 800
const DISTANCE = pb.Distance_Cosine
const THRESHOLD_SCORE float32 = 0.9
var Indexes = []string{"username", "email"}

// Benchmarking parameters
const BENCHMARK_ITERATIONS = 100000
const BATCH_SIZE = 10000

