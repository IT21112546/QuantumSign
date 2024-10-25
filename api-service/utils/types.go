package utils

// User register request body
type CreateUserBody struct {
	PubKey   []float32
	Username string
	Email    string
}

// User search response body
type UserSearchResult struct {
	Username string
	Email    string
}
