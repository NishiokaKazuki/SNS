package api

import (
	"context"
	"log"
	"net"
	"server/generated/enums"
	"server/generated/messages"
	pb "server/generated/services"
	"server/model/db"
	"server/model/tables"
	qr "server/queries"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) Auth(ctx context.Context, in *messages.AuthRequest) (*messages.AuthResponse, error) {

	return &messages.AuthResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Token:      "test",
	}, nil
}

func (s *server) User(ctx context.Context, in *messages.UserRequest) (*messages.UserResponse, error) {
	var (
		user tables.AppUsers
	)

	user, err := qr.GetUser(db.GetDBConnect(), ctx, 1)
	if err != nil {
		return &messages.UserResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
			User:       &messages.UserResponse_AppUser{},
		}, status.Error(codes.NotFound, err.Error())
	}

	return &messages.UserResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		User: &messages.UserResponse_AppUser{
			Handle:    user.Handle,
			Name:      user.Name,
			Birthday:  user.Birthday.String(),
			Profile:   user.Profile,
			IsPrivate: user.IsPrivate,
		},
	}, status.Error(codes.OK, "")
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
