package interceptor

import (
	"context"

	"server/auth"
	"server/model/db"
	"server/queries"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	PermissionAuth   = "Auth"
	PermissionSignIn = "SignIn"
	PermissionSignUp = "SignUp"
	PermissionUser   = "User"
)

var routes = map[string]string{
	"/pb.Service/Auth":   PermissionAuth,
	"/pb.Service/SignIn": PermissionSignIn,
	"/pb.Service/SignUp": PermissionSignUp,
	"/pb.Service/User":   PermissionUser,
}

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// info.FullMethodにメソッドのフルパスが入っている。
		users, _ := queries.GetUserByToken(db.GetDBConnect(), ctx, auth.GetToken(ctx).Subject)

		if users.Id == 0 {
			switch routes[info.FullMethod] {
			case PermissionAuth:
				return handler(ctx, req)
			case PermissionSignIn:
				return handler(ctx, req)
			case PermissionSignUp:
				return handler(ctx, req)
			default:
				return nil, status.Error(
					codes.PermissionDenied,
					"Please signIn or signUp",
				)
			}
		} else {
			switch routes[info.FullMethod] {
			case PermissionAuth:
				return handler(ctx, req)
			case PermissionUser:
				return handler(ctx, req)
			default:
				return nil, status.Error(
					codes.NotFound,
					"Cannot access"+info.FullMethod,
				)
			}
		}
	}
}
