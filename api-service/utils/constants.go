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

// MongoDB database parameters
const MONGO_ADDR = "mongodb://root:sliit123@localhost:27017"
const MONGO_DATABASE = "SSO_DB"
const MONGO_COLLECTION = "clients"

// Benchmarking parameters
const BENCHMARK_ITERATIONS = 100000
const BATCH_SIZE = 10000

// Error & success messages
const USER_ALREADY_EXISTS_MSG = "User with public key already exists"
const USER_CREATED_MSG = "User created successfully"
const INTERNAL_SERVER_ERROR_MSG = "Internal server error"
const BASE64_DECODING_ERROR_MSG = "Error decoding base64 encoded string"
const ERROR_UNMARSHALLING_MSG = "Error unmarshalling public key"

// Response codes
var SUCCESS int = 200
var BAD_REQUEST int = 400
var INTERNAL_SERVER_ERROR int = 500
var CONFLICT int = 409
var CREATED int = 201

// SSO Default Redirect URL
var DEFAULT_REDIRECT_URL = "http://sso.qsign.io/redirect"
