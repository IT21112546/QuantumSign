package utils

import (
	"fmt"
	"regexp"
	"net/http"
	"encoding/json"
)

// To check for errors
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Validate if emails are valid
func ValidateEmail(email string) bool {
	// Simple email validation regex
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}


// Respond to HTTP errors in JSON format
func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]interface{}{
		"error": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
