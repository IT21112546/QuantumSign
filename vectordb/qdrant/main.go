package main

import (
	"fmt"
	"path/filepath"

	"gitlab.sliit.lk/r24-055/r24-055/vectordb/utils"
)

func main() {

	// Create a new collection and index
	utils.NewQdrantCollection()
	utils.NewKeywordIndex()

	fmt.Printf("What operation do you want to do?\n")
	fmt.Printf("1: Register\n")
	fmt.Printf("2: Login\n")
	fmt.Printf("3: Benchmark\n\n")
	fmt.Printf("Choice: ")

	// Take input
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		register()

	case 2:
		login()

	case 3:
		register()
	}
}

func register() {

	var keyPath string
	fmt.Printf("Enter the path to the public key file: ")
	fmt.Scanln(&keyPath)

	var username string
	fmt.Printf("Enter the username: ")
	fmt.Scanln(&username)

	var email string
	fmt.Printf("Enter the email: ")
	fmt.Scanln(&email)

	// Import public key (from current working directory)
	path, err := filepath.Abs(keyPath)
	pubVector := utils.ImportPubKey(path)

	// Create a user
	pointsStruct, err := utils.CreateUserPoints(utils.CreateUserBody{
		PubKey:   pubVector,
		Username: username,
		Email:    email,
	})

	utils.CheckErr(err)

	// Add user to the collection
	err, insertTime := utils.UpsertPoints(pointsStruct)
	utils.CheckErr(err)
	fmt.Printf("Data inserted in %.3f ms\n", insertTime)
}

func login() {
	var keyPath string
	fmt.Printf("Enter the path to the public key file: ")
	fmt.Scanln(&keyPath)

	// Import public key (from current working directory)
	path, err := filepath.Abs(keyPath)
	utils.CheckErr(err)
	pubVector := utils.ImportPubKey(path)

	// Search for the user
	res, isRegistered, _ := utils.GetUserData(pubVector)

	if isRegistered {
		fmt.Printf("Hello %s. Your email is: %s\n", res.Username, res.Email)
	} else {
		fmt.Printf("User not registered\n")
	}
}

func benchmark() {
	pubKeys := utils.GenPubKeys(utils.BENCHMARK_ITERATIONS)
	utils.BenchmarkAverages(pubKeys, utils.BATCH_SIZE)
}
