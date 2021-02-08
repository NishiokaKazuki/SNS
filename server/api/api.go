package api

import (
	"context"
	"io"
	"log"
	"net"
	"server/auth"
	"server/generated/enums"
	"server/generated/messages"
	"server/generated/services"
	pb "server/generated/services"
	"server/interceptor"
	"server/model/db"
	"server/model/tables"
	qr "server/queries"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) SignIn(ctx context.Context, in *messages.SignInRequest) (*messages.SignInResponse, error) {

	user, err := qr.GetUserByPass(db.GetDBConnect(), ctx, in.Handle, auth.HashPw(in.Password))
	if err != nil || user.Id == 0 {
		return &messages.SignInResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unauthenticated, err.Error())
	}

	token := auth.CreateToken(user)

	affected, err := qr.InsertTokens(db.GetDBConnect(), ctx, tables.Tokens{
		UserId: user.Id,
		Token:  token,
	})
	if err != nil {
		return &messages.SignInResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unauthenticated, err.Error())
	}
	if affected != true {
		// wip
	}

	return &messages.SignInResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Token:      token,
	}, nil
}

func (s *server) SignUp(ctx context.Context, in *messages.SignUpRequest) (*messages.SignUpResponse, error) {

	user, err := qr.GetUserByHandle(db.GetDBConnect(), ctx, in.Handle)
	if err != nil {
		return &messages.SignUpResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unauthenticated, err.Error())
	}
	if user.Id == 0 {
		return &messages.SignUpResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.AlreadyExists, "Handle Already Exists")
	}

	// wip encode
	affected, err := qr.InsertAppUser(db.GetDBConnect(), ctx, tables.AppUsers{
		Handle:   in.Handle,
		Password: auth.HashPw(in.Password),
		Name:     in.Name,
	})
	if err != nil {
		return &messages.SignUpResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unauthenticated, err.Error())
	}
	if affected != true {
		// wip
	}

	user, err = qr.GetUserByPass(db.GetDBConnect(), ctx, in.Handle, auth.HashPw(in.Password))
	if err != nil || user.Id == 0 {
		return &messages.SignUpResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unauthenticated, err.Error())
	}

	token := auth.CreateToken(user)

	affected, err = qr.InsertTokens(db.GetDBConnect(), ctx, tables.Tokens{
		UserId: user.Id,
		Token:  token,
	})
	if err != nil {
		return &messages.SignUpResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unauthenticated, err.Error())
	}
	if affected != true {
		// wip
	}

	return &messages.SignUpResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Token:      token,
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

func (s *server) Message(stream services.Service_MessageServer) error {
	ctx := stream.Context()
	token, _ := auth.GetToken(ctx)
	users, _ := qr.GetUserByToken(db.GetDBConnect(), ctx, token)

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}

		if in.GetBody() == "" {
			continue
		}
		log.Println(in)
		_, err = qr.InsertMessageLogs(ctx, db.GetDBConnect(), tables.MessageLogs{
			UserId:  users.Id,
			IsGroup: !in.GetIsUser(),
			Body:    in.GetBody(),
		})

		messageLog, err := qr.GetMessageLogs(ctx, db.GetDBConnect(), tables.MessageLogs{
			UserId: users.Id,
		})

		var userIds []uint64
		if !in.GetIsUser() {
			qr.InsertLogToGroup(ctx, db.GetDBConnect(), tables.LogToGroups{
				GroupId: in.GetSendId(),
				LogId:   messageLog.Id,
			})
			users, err := qr.FindUsersByGroupId(ctx, db.GetDBConnect(), in.GetSendId())
			if err != nil {
				break
			}
			for _, u := range users {
				userIds = append(userIds, u.Id)
			}
		} else {
			userIds = append(userIds, in.GetSendId())
		}

		for _, id := range userIds {
			qr.InsertLogToUsers(ctx, db.GetDBConnect(), tables.LogToUsers{
				UserId:      id,
				LogId:       messageLog.Id,
				IsConfirmed: false,
			})
		}
	}

	return nil
}

func (s *server) Receive(req *messages.ReceivedRequest, stream pb.Service_ReceiveServer) error {
	ctx := stream.Context()
	token, _ := auth.GetToken(ctx)
	users, _ := qr.GetUserByToken(db.GetDBConnect(), ctx, token)

	for {
		var logIds []uint64
		messageLogs, err := qr.FindMessageLogsByUserId(ctx, db.GetDBConnect(), users.Id)
		if err != nil {
			log.Println(err)
			break
		}
		for _, log := range messageLogs {
			var id uint64
			if log.IsGroup {
				logToGroup, _ := qr.GetLogToGroup(ctx, db.GetDBConnect(), log.Id)
				id = logToGroup.GroupId
			} else {
				id = log.UserId
			}

			err = stream.Send(&messages.ReceivedResponse{
				Status:     true,
				StatusCode: enums.StatusCodes_SUCCESS,
				Body:       log.Body,
				IsUser:     !log.IsGroup,
				SendId:     id,
			})
			if err != nil {
				break
			}
			logIds = append(logIds, log.Id)
		}
		if err != nil {
			break
		}

		if len(messageLogs) > 0 {
			_, err = qr.UpdatelogToUsers(ctx, db.GetDBConnect(), tables.LogToUsers{
				IsConfirmed: true,
			}, logIds, users.Id)
		}
		time.Sleep(time.Second * 1)
	}

	return nil
}

func ListenAndServe(port string) {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.Auth),
			interceptor.AuthorizationUnaryServerInterceptor(),
		),
		),
		grpc.ChainStreamInterceptor(
			grpc_auth.StreamServerInterceptor(auth.StreamServerAuthorized),
		),
	)
	pb.RegisterServiceServer(s, &server{})

	log.Println("starting server", port)
	err = s.Serve(listenPort)
	if err != nil {
		log.Fatal("failed open", port)
	}
}
