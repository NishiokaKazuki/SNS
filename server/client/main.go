package main

import (
	"context"
	"log"
	"os"
	"time"

	"server/generated/messages"
	pb "server/generated/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	host = "localhost:49200"
)

func main() {

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed connect: %s", err)
	}

	defer conn.Close()

	client := pb.NewServiceClient(conn)
	// タイムアウトを20秒に設定する
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	sw := os.Args
	switch sw[1] {
	case "user":
		user(ctx, client)

	case "signIn":
		signIn(ctx, client)
	default:
		log.Println("not args")

	}
}

func user(ctx context.Context, client pb.ServiceClient) {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("no more args")
	}
	token := args[2]
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.User(ctx, &messages.UserRequest{
		Token: "not used",
	})

	if err == nil {
		log.Println(res)
	} else {
		log.Println(err)
	}
}

func signIn(ctx context.Context, client pb.ServiceClient) {
	args := os.Args
	if len(args) < 3 {
		log.Fatalln("no more args")
	}
	md := metadata.New(map[string]string{"authorization": "Bearer testtoken"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	handle := args[2]
	password := args[3]

	res, err := client.SignIn(ctx, &messages.SignInRequest{
		Handle:   handle,
		Password: password,
	})

	if err == nil {
		log.Println(res)
		// log.Printf("%#v\n", res.Company.Name)
	} else {
		log.Println(err)
	}
}
