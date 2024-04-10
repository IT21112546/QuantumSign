package algorithms

import "os"

func ExportJWT(token string, filename string) {
	// Export token to file
	tokenFile, err := os.Create("keys/jwt/" + filename)
	if err != nil {
		panic(err)
	}
	defer tokenFile.Close()
	_, err = tokenFile.Write([]byte(token))
	if err != nil {
		panic(err)
	}
}

func ImportJWT(filename string) (string, error) {
	// Read the token bytes from file
	tokenBytes, err := os.ReadFile("keys/jwt/" + filename)
	if err != nil {
		return "", err
	}

	return string(tokenBytes), nil
}
