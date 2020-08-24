package main

import (
	"context"
	"log"
	"os"
	"time"

	"server/generated/messages"
	pb "server/generated/services"

	"google.golang.org/grpc"
)

const (
	host = "localhost:49200"
)

func main() {

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failrd connect: %s", err)
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

	}
}

func user(ctx context.Context, client pb.ServiceClient) {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("no more args")
	}
	// id, _ := strconv.Atoi(args[2])

	res, err := client.Auth(ctx, &messages.AuthRequest{
		Token: "test",
	})
	if err == nil {
		log.Println(res)
		// log.Printf("%#v\n", res.Company.Name)
	} else {
		log.Println(err)
	}
}
