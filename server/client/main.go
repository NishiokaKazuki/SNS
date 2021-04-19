package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"server/generated/messages"
	pb "server/generated/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	host = "localhost:80"
)

func main() {

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed connect: %s", err)
	}

	defer conn.Close()

	client := pb.NewServiceClient(conn)
	// タイムアウトを20秒に設定する
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*300)
	defer cancel()

	sw := os.Args
	switch sw[1] {
	case "user":
		user(ctx, client)

	case "signIn":
		signIn(ctx, client)

	case "mes":
		finish := make(chan bool)

		go message(ctx, client)
		go receive(ctx, client)
		<-finish
		message(ctx, client)
	case "rec":
		receive(ctx, client)
	case "createGroup":
		createGroup(ctx, client)
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
	res, err := client.User(ctx, &messages.UserRequest{})

	if err == nil {
		log.Println(res)
	} else {
		log.Println(err)
	}
}

func createGroup(ctx context.Context, client pb.ServiceClient) {
	args := os.Args
	if len(args) < 3 {
		log.Fatalln("no more args")
	}
	token := args[2]
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.CreateGroup(ctx, &messages.CreateGroupRequest{
		Token: token,
		Group: &messages.UserGroup{
			Name: args[3],
		},
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

func message(ctx context.Context, client pb.ServiceClient) {
	args := os.Args
	if len(args) < 3 {
		log.Fatalln("no more args")
	}
	token := args[2]
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	for {
		var send uint64
		var body string
		fmt.Scan(&send)
		fmt.Scan(&body)
		res, err := client.Message(ctx, &messages.MessageRequest{
			IsUser: true,
			SendId: send,
			Body:   body,
		})
		if err != nil {
			log.Println(res)
			log.Println(err)
		}
	}
}

func receive(ctx context.Context, client pb.ServiceClient) {
	args := os.Args
	if len(args) < 3 {
		log.Fatalln("no more args")
	}
	token := args[2]
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	stream, err := client.Receive(ctx, &messages.ReceivedRequest{})
	if err != nil {
		log.Println(err)
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		// お返し
		log.Println(in.GetBody())
		log.Println(in.GetIsUser())
		log.Println(in.GetSendId())
	}
}
