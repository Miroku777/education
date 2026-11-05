package main

// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// linux: sudo apt install -y protobuf-compiler		- protoc

// go get google.golang.org/grpc
// go get google.golang.org/protobuf

// protoc --go_out=. --go-grpc_out=. user.proto 	- сбор
import (
	"context"
	"fmt"
	"lesson20/user"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	client := user.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	authResp, err := client.CheckUserCredo(ctx, &user.AuthRequest{
		Email:    "ivanov@example.com",
		Password: "123",
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(authResp.Answer)
}
