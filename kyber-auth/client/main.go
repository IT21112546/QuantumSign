package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/circl/kem/kyber/kyber512"
	"gitlab.sliit.lk/r24-055/r24-055/kyber-auth"
	pb "gitlab.sliit.lk/r24-055/r24-055/kyber-auth/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverPublicKey = utils.ImportPubKey("server_key.pub")
const BENCHMARK_ITER = 100

type User struct {
	PublicKeyPath string
	Username      string
	Email         string
}

func register(user User, conn *grpc.ClientConn) {
	c := pb.NewRegisterServiceClient(conn)
	clientPublicKey, err := utils.ImportPubKey(user.PublicKeyPath).MarshalBinary()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Register(ctx, &pb.RegisterRequest{
		PublicKey: clientPublicKey,
		Username: user.Username,
		Email: user.Email,
	})
	utils.CheckErr(err)

	// Decrypt received data
	registerSuccess := r.Success
	if registerSuccess {
		fmt.Println("You are registered")
	} else {
		fmt.Println("Registration failed")
	}
}

func login(pkPath string, conn *grpc.ClientConn, isBenchmark bool) {

	c := pb.NewKeyExchangeServiceClient(conn)
	clientPublicKey, err := utils.ImportPubKey(pkPath).MarshalBinary()
	encryptedRandomString, randomString, err := kyber512.Scheme().Encapsulate(serverPublicKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.KeyExchange(ctx, &pb.KeyExchangeRequest{
		PublicKey: clientPublicKey,
		Kem:       encryptedRandomString,
	})
	utils.CheckErr(err)

	// Decrypt received data
	decryptedSecret := utils.CyclicXOR(randomString, r.EncryptedSharedSecret)
	if !isBenchmark {
		fmt.Println("You are logged in")
		fmt.Printf("Access Token: %s\n", string(decryptedSecret))
	}
}

func benchmark(iter uint8, keyPath string, conn *grpc.ClientConn) {
	totalTime := time.Duration(0)

	for i := uint8(0); i < iter; i++ {
		start := time.Now()

		login(keyPath, conn, true)

		elapsed := time.Since(start)
		fmt.Printf("Took %v for iteration %d\n", elapsed, i)
		totalTime += elapsed
	}

	fmt.Printf("\nAverage time taken: %v\n", totalTime/time.Duration(iter))
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	utils.CheckErr(err)

	utils.CheckErr(err)

	fmt.Printf("What operation do you want to do?\n")
	fmt.Printf("1: Register\n")
	fmt.Printf("2: Login\n")
	fmt.Printf("3: Benchmark\n")
	fmt.Printf("Choice: ")

	// Take input
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		var keyPath string
		fmt.Printf("Enter the path to your public key: ")
		fmt.Scanln(&keyPath)

		var username string
		fmt.Printf("Enter your username: ")
		fmt.Scanln(&username)

		var email string
		fmt.Printf("Enter your email: ")
		fmt.Scanln(&email)

		user := User{
			PublicKeyPath: keyPath,
			Username:      username,
			Email:         email,
		}

		register(user, conn)

	case 2:
		var keyPath string
		fmt.Printf("Enter the path to the public key file: ")
		fmt.Scanln(&keyPath)

		login(keyPath, conn, false)

	case 3:
		var keyPath string
		fmt.Printf("Enter the path to the public key file: ")
		fmt.Scanln(&keyPath)

		benchmark(BENCHMARK_ITER, keyPath, conn)
	}

}
