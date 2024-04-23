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

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	utils.CheckErr(err)
	defer conn.Close()
	c := pb.NewKeyExchangeServiceClient(conn)

	clientPublicKey, err := utils.ImportPubKey("client_key.pub").MarshalBinary()

	serverPublicKey := utils.ImportPubKey("server_key.pub")
	encryptedRandomString, randomString, err := kyber512.Scheme().Encapsulate(serverPublicKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.KeyExchange(ctx, &pb.KeyExchangeRequest{
		PublicKey: clientPublicKey,
		Kem: encryptedRandomString,
	})
	utils.CheckErr(err)

	// Decrypt received data
	decryptedSecret := utils.CyclicXOR(randomString, r.EncryptedSharedSecret)
	fmt.Printf("Decrypted result: %s\n", string(decryptedSecret))
}
