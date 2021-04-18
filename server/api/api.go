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

func (s *server) CreateGroup(ctx context.Context, in *messages.CreateGroupRequest) (*messages.CreateGroupResponse, error) {
	token, _ := auth.GetToken(ctx)
	user, _ := qr.GetUserByToken(db.GetDBConnect(), ctx, token)
	affected, err := qr.InsertUserGroup(ctx, db.GetDBConnect(), tables.UserGroups{
		Name: in.GetGroup().Name,
	})
	if err != nil {
		log.Println(affected)
		return &messages.CreateGroupResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
			Group:      &messages.UserGroup{},
		}, status.Error(codes.ResourceExhausted, err.Error())
	}

	userGroup, _ := qr.GetUserGroupByName(ctx, db.GetDBConnect(), in.GetGroup().Name)
	affected, err = qr.InsertGroupToUser(ctx, db.GetDBConnect(), tables.GroupToUsers{
		GroupId: userGroup.Id,
		UserId:  user.Id,
	})
	if err != nil {
		log.Println(affected)
		return &messages.CreateGroupResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
			Group:      &messages.UserGroup{},
		}, status.Error(codes.ResourceExhausted, err.Error())
	}

	return &messages.CreateGroupResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Group:      &messages.UserGroup{},
	}, nil
}

func (s *server) Followers(ctx context.Context, in *messages.FollowersRequest) (*messages.FollowersResponse, error) {
	var appUsers []*messages.AppUser
	token, _ := auth.GetToken(ctx)
	user, _ := qr.GetUserByToken(db.GetDBConnect(), ctx, token)

	users, err := qr.FindAppUsersByToFollows(ctx, db.GetDBConnect(), user.Id)
	if err != nil {
		return &messages.FollowersResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
		}, status.Error(codes.ResourceExhausted, err.Error())
	}

	for _, u := range users {
		appUsers = append(appUsers, &messages.AppUser{
			Id:        u.Id,
			Handle:    u.Handle,
			Name:      u.Name,
			Birthday:  u.Birthday.Format("2006/01/02"),
			Profile:   u.Profile,
			IsPrivate: u.IsPrivate,
		})
	}
	return &messages.FollowersResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		AppUsers:   appUsers,
	}, nil
}

func (s *server) GetGroups(ctx context.Context, in *messages.GetGroupsRequest) (*messages.GetGroupsResponse, error) {
	var groups []*messages.UserGroup

	token, _ := auth.GetToken(ctx)
	user, _ := qr.GetUserByToken(db.GetDBConnect(), ctx, token)
	userGroups, err := qr.FindUserGroupsByUserId(ctx, db.GetDBConnect(), user.Id)
	if err != nil {
		return &messages.GetGroupsResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unavailable, err.Error())
	}

	for _, g := range userGroups {
		var (
			groupUsers  []*messages.AppUser
			inviteUsers []*messages.AppUser
		)

		iu, err := qr.FindAppUsersByInvitesGroupId(ctx, db.GetDBConnect(), g.Id)
		if err != nil {
			return &messages.GetGroupsResponse{
				Status:     false,
				StatusCode: enums.StatusCodes_FAILED_AUTH,
			}, status.Error(codes.Unavailable, err.Error())
		}
		au, err := qr.FindAppUsersByGroupId(ctx, db.GetDBConnect(), g.Id)
		if err != nil {
			return &messages.GetGroupsResponse{
				Status:     false,
				StatusCode: enums.StatusCodes_FAILED_AUTH,
			}, status.Error(codes.Unavailable, err.Error())
		}
		for _, i := range iu {
			inviteUsers = append(inviteUsers, &messages.AppUser{
				Id:        i.Id,
				Handle:    i.Handle,
				Name:      i.Name,
				Birthday:  i.Birthday.String(),
				Profile:   i.Profile,
				IsPrivate: i.IsPrivate,
			})
		}
		for _, i := range au {
			groupUsers = append(inviteUsers, &messages.AppUser{
				Id:        i.Id,
				Handle:    i.Handle,
				Name:      i.Name,
				Birthday:  i.Birthday.String(),
				Profile:   i.Profile,
				IsPrivate: i.IsPrivate,
			})
		}

		groups = append(groups, &messages.UserGroup{
			Id:           g.Id,
			Name:         g.Name,
			Users:        groupUsers,
			InvitedUsers: inviteUsers,
		})
	}

	return &messages.GetGroupsResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Groups:     groups,
	}, nil
}

func (s *server) GetInvitedGroups(ctx context.Context, in *messages.GetInvitedGroupsRequest) (*messages.GetInvitedGroupsResponse, error) {
	var groups []*messages.UserGroup

	token, _ := auth.GetToken(ctx)
	user, _ := qr.GetUserByToken(db.GetDBConnect(), ctx, token)
	userGroups, err := qr.FindInviteUserToGroupsByUserId(ctx, db.GetDBConnect(), user.Id)
	if err != nil {
		return &messages.GetInvitedGroupsResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED_AUTH,
		}, status.Error(codes.Unavailable, err.Error())
	}

	for _, g := range userGroups {
		var (
			groupUsers  []*messages.AppUser
			inviteUsers []*messages.AppUser
		)

		iu, err := qr.FindAppUsersByInvitesGroupId(ctx, db.GetDBConnect(), g.Id)
		if err != nil {
			return &messages.GetInvitedGroupsResponse{
				Status:     false,
				StatusCode: enums.StatusCodes_FAILED_AUTH,
			}, status.Error(codes.Unavailable, err.Error())
		}
		au, err := qr.FindAppUsersByGroupId(ctx, db.GetDBConnect(), g.Id)
		if err != nil {
			return &messages.GetInvitedGroupsResponse{
				Status:     false,
				StatusCode: enums.StatusCodes_FAILED_AUTH,
			}, status.Error(codes.Unavailable, err.Error())
		}
		for _, i := range iu {
			inviteUsers = append(inviteUsers, &messages.AppUser{
				Id:        i.Id,
				Handle:    i.Handle,
				Name:      i.Name,
				Birthday:  i.Birthday.String(),
				Profile:   i.Profile,
				IsPrivate: i.IsPrivate,
			})
		}
		for _, i := range au {
			groupUsers = append(inviteUsers, &messages.AppUser{
				Id:        i.Id,
				Handle:    i.Handle,
				Name:      i.Name,
				Birthday:  i.Birthday.String(),
				Profile:   i.Profile,
				IsPrivate: i.IsPrivate,
			})
		}

		groups = append(groups, &messages.UserGroup{
			Id:           g.Id,
			Name:         g.Name,
			Users:        groupUsers,
			InvitedUsers: inviteUsers,
		})
	}

	return &messages.GetInvitedGroupsResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Groups:     groups,
	}, nil
}

func (s *server) JoinInvitedGroups(ctx context.Context, in *messages.JoinInvitedGroupsRequest) (*messages.JoinInvitedGroupsResponse, error) {
	token, _ := auth.GetToken(ctx)
	user, err := qr.GetUserByToken(db.GetDBConnect(), ctx, token)
	if err != nil {
		return &messages.JoinInvitedGroupsResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
		}, status.Error(codes.NotFound, err.Error())
	}

	group, err := qr.GetUserGroupByInviteUserToGroup(ctx, db.GetDBConnect(), tables.InviteUserToGroups{
		UserId:  user.Id,
		GroupId: in.GetGroupId(),
	})
	if err != nil {
		return &messages.JoinInvitedGroupsResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
		}, status.Error(codes.PermissionDenied, err.Error())
	}
	if group.Id != in.GetGroupId() {
		return &messages.JoinInvitedGroupsResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
		}, status.Error(codes.PermissionDenied, err.Error())
	}

	affected, err := qr.DeleteInviteUserToGroupByUserId(ctx, db.GetDBConnect(), tables.InviteUserToGroups{
		UserId:  user.Id,
		GroupId: in.GetGroupId(),
	})
	if affected != true {
		return &messages.JoinInvitedGroupsResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
		}, status.Error(codes.PermissionDenied, err.Error())
	}
	if err != nil {
		return &messages.JoinInvitedGroupsResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
		}, status.Error(codes.PermissionDenied, err.Error())
	}

	qr.InsertGroupToUser(ctx, db.GetDBConnect(), tables.GroupToUsers{
		UserId:  user.Id,
		GroupId: in.GetGroupId(),
	})

	return &messages.JoinInvitedGroupsResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
	}, nil
}

func (s *server) InviteToGroup(ctx context.Context, in *messages.InviteToGroupRequest) (*messages.InviteToGroupResponse, error) {
	qr.InsertInviteUserToGroup(ctx, db.GetDBConnect(), tables.InviteUserToGroups{
		GroupId: in.GetGroupId(),
		UserId:  in.GetUserId(),
	})

	return &messages.InviteToGroupResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
	}, nil
}

func (s *server) User(ctx context.Context, in *messages.UserRequest) (*messages.UserResponse, error) {
	token, _ := auth.GetToken(ctx)
	user, err := qr.GetUserByToken(db.GetDBConnect(), ctx, token)
	if err != nil {
		return &messages.UserResponse{
			Status:     false,
			StatusCode: enums.StatusCodes_FAILED,
			User:       &messages.AppUser{},
		}, status.Error(codes.NotFound, err.Error())
	}

	return &messages.UserResponse{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		User: &messages.AppUser{
			Id:        user.Id,
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
			users, err := qr.FindAppUsersByGroupId(ctx, db.GetDBConnect(), in.GetSendId())
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
	// grpcWebServer := grpcweb.WrapServer(
	// 	s,
	// 	grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	// )
	// httpServer := &http.Server{
	// 	Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		if r.ProtoMajor == 2 {
	// 			grpcWebServer.ServeHTTP(w, r)
	// 		} else {
	// 			w.Header().Set("Access-Control-Allow-Origin", "*")
	// 			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
	// 			w.Header().Set("grpc-status", "")
	// 			w.Header().Set("grpc-message", "")
	// 			if grpcWebServer.IsGrpcWebRequest(r) {
	// 				grpcWebServer.ServeHTTP(w, r)
	// 			}
	// 		}
	// 	}), &http2.Server{}),
	// }

	// wrappedServer := grpcweb.WrapServer(
	// 	s,
	// 	grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	// )
	// http.Handle("/", wrappedServer)

	log.Println("starting server", port)
	err = s.Serve(listenPort)
	// httpServer.Serve(listenPort)
	if err != nil {
		log.Fatal("failed open", port)
	}
}
