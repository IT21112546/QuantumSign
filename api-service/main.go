package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cloudflare/circl/kem/kyber/kyber512"
	"github.com/cloudflare/circl/sign/dilithium"
	"github.com/google/uuid"
	algorithms "github.com/meeranh/register_service/base"
	"github.com/meeranh/register_service/utils"
	"gitlab.sliit.lk/r24-055/r24-055/dilithium-jwt/base/jwt/v5"
)

// Register request body structure
type RegisterRequest struct {
	KyberPublicKey string `json:"kyberPublicKey"`
	Username       string `json:"username"`
	Email          string `json:"email"`
}

// Login request body structure (Kyber)
type LoginRequest struct {
	KyberPublicKey string `json:"kyberPublicKey"`
	ClientId       string `json:"clientId"`
}

// JWT Validate request body structure
type ValidateRequest struct {
	JWTToken string `json:"jwtToken"`
	ClientId string `json:"clientId"`
}

// Load Dilithium public and private keys
var pubKey, priKey, err = algorithms.ImportDilithiumKeys()

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse the JSON request body
	var req RegisterRequest
	if err := json.Unmarshal(body, &req); err != nil {
		utils.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate the provided email
	if !utils.ValidateEmail(req.Email) {
		utils.ErrorResponse(w, "Invalid email", http.StatusBadRequest)
		return
	}

	// Load the public key
	publicKeyBytes, err := utils.LoadPubKey(req.KyberPublicKey)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resMsg, resCode := utils.CreateUser(utils.CreateUserBody{
		PubKey:   publicKeyBytes,
		Username: req.Username,
		Email:    req.Email,
	})

	// Create JSON response
	response := map[string]interface{}{
		"message": resMsg,
	}

	// Marshal the response to JSON and write it
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.ErrorResponse(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resCode)
	w.Write(jsonResponse)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse the JSON request body
	var req LoginRequest
	if err := json.Unmarshal(body, &req); err != nil {
		utils.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Load the public key
	publicKeyBytes, err := utils.LoadPubKey(req.KyberPublicKey)
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch the username and email from the collection
	dbResponse, isRegistered, _ := utils.GetUserData(publicKeyBytes)

	// If the user is not registered
	if !isRegistered {
		utils.ErrorResponse(w, "User is not registered", http.StatusNotFound)
		return
	}

	// If clientId is provided, use the client's private key and redirectUrl
	var signingKey dilithium.PrivateKey
	var redirectUrl string

	if req.ClientId != "" {
		clientData, err := utils.GetClientData(req.ClientId)
		if err != nil {
			utils.ErrorResponse(w, "Client not found", http.StatusNotFound)
			return
		}

		privateKeyBytes, err := base64.StdEncoding.DecodeString(clientData["privateKey"])
		if err != nil {
			utils.ErrorResponse(w, "Invalid private key", http.StatusInternalServerError)
			return
		}

		signingKey = dilithium.Mode3.PrivateKeyFromBytes(privateKeyBytes)
		redirectUrl = clientData["redirectUrl"]
	} else {
		signingKey = priKey
		redirectUrl = utils.DEFAULT_REDIRECT_URL
	}

	// Generate JWT token
	token, err := algorithms.GenTokenDilithium(jwt.MapClaims{
		"username": dbResponse.Username,
		"email":    dbResponse.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}, signingKey)

	if err != nil {
		utils.ErrorResponse(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Write the token
	jsonResponse, err := json.Marshal(map[string]interface{}{
		"accessToken": token,
		"redirectUrl": redirectUrl,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func generateDilithiumKeyPairHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Generate a Dilithium keypair
	publicKey, privateKey := algorithms.GenDilithium()

	// Encode the keys to base64
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKey.Bytes())
	encodedPrivateKey := base64.StdEncoding.EncodeToString(privateKey.Bytes())

	response := map[string]interface{}{
		"publicKey":  encodedPublicKey,
		"privateKey": encodedPrivateKey,
	}

	// Marshal the response to JSON and write it
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.ErrorResponse(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func generateKyberKeyPairHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Generate a Kyber keypair
	publicKey, privateKey, err := kyber512.Scheme().GenerateKeyPair()

	// If there is an error
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the keys into bytes
	publicKeyBytes, err := publicKey.MarshalBinary()
	privateKeyBytes, err := privateKey.MarshalBinary()

	// If there is an error
	if err != nil {
		utils.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the keys to base64
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKeyBytes)
	encodedPrivateKey := base64.StdEncoding.EncodeToString(privateKeyBytes)

	response := map[string]interface{}{
		"publicKey":  encodedPublicKey,
		"privateKey": encodedPrivateKey,
	}

	// Marshal the response to JSON and write it
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.ErrorResponse(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse the JSON request body
	var req ValidateRequest
	if err := json.Unmarshal(body, &req); err != nil {
		utils.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var publicKeyBytes []byte
	if req.ClientId != "" {
		clientData, err := utils.GetClientData(req.ClientId)
		if err != nil {
			utils.ErrorResponse(w, "Client not found", http.StatusNotFound)
			return
		}
		publicKeyBytes, err = base64.StdEncoding.DecodeString(clientData["publicKey"])
		if err != nil {
			utils.ErrorResponse(w, "Invalid public key", http.StatusInternalServerError)
			return
		}
	} else {
		publicKeyBytes = pubKey.Bytes()
		if err != nil {
			utils.ErrorResponse(w, "Invalid public key", http.StatusInternalServerError)
			return
		}
	}

	dilithiumPubKey := dilithium.Mode3.PublicKeyFromBytes(publicKeyBytes)

	// Validate the JWT token
	isExpired, isValid, err := algorithms.VerifyTokenDilithium(req.JWTToken, dilithiumPubKey)
	if err != nil {
		utils.ErrorResponse(w, "The token signature appears to be invalid, or might belong to another client", http.StatusInternalServerError)
		return
	}

	// Create JSON response
	response := map[string]interface{}{
		"isValid":   isValid,
		"isExpired": isExpired,
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func clientRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorResponse(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req map[string]string
	if err := json.Unmarshal(body, &req); err != nil {
		utils.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	clientName := req["clientName"]
	redirectUrl := req["redirectUrl"]

	// Generate clientId as UUID
	clientId := uuid.NewString()

	// Generate Dilithium key pair
	publicKey, privateKey := algorithms.GenDilithium()
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKey.Bytes())
	encodedPrivateKey := base64.StdEncoding.EncodeToString(privateKey.Bytes())

	// Save the data into MongoDB
	clientData := map[string]string{
		"clientName":  clientName,
		"clientId":    clientId,
		"publicKey":   encodedPublicKey,
		"privateKey":  encodedPrivateKey,
		"redirectUrl": redirectUrl,
	}

	if err := utils.SaveClientData(clientData); err != nil {
		utils.ErrorResponse(w, "Failed to save client data", http.StatusInternalServerError)
		return
	}

	// Respond with the clientId
	response := map[string]interface{}{
		"clientId":   clientId,
		"publicKey":  encodedPublicKey,
		"privateKey": encodedPrivateKey,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	// Initiate MongoDB connection
	utils.InitMongoDB()

	// Create a new collection and index
	utils.NewQdrantCollection()
	utils.NewKeywordIndex()

	// Declare the API endpoints
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/generate/dilithium", generateDilithiumKeyPairHandler)
	http.HandleFunc("/generate/kyber", generateKyberKeyPairHandler)
	http.HandleFunc("/client-registration", clientRegistrationHandler)

	// Start the server
	fmt.Println("Server is listening on port 4444...")
	err = http.ListenAndServe(":4444", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
