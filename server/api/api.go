package api

import (
	"context"
	"log"
	"net"
	"server/generated/enums"
	"server/generated/messages"
	pb "server/generated/services"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Auth(ctx context.Context, in *messages.AuthRequest) (*messages.AuthResponse, error) {

	return &messages.AuthResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Token:      "test",
	}, nil
}

func ListenAndServe(port string) {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &server{})

	log.Println("starting server", port)
	err = s.Serve(listenPort)
	if err != nil {
		log.Fatal("failed open", port)
	}
}
