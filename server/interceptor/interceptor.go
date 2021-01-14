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
	"/services.Service/Auth":   PermissionAuth,
	"/services.Service/SignIn": PermissionSignIn,
	"/services.Service/SignUp": PermissionSignUp,
	"/services.Service/User":   PermissionUser,
}

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if routes[info.FullMethod] == PermissionSignIn ||
			routes[info.FullMethod] == PermissionSignUp {
			return handler(ctx, req)
		}

		token, err := auth.GetToken(ctx)
		if err != nil {
			return nil, status.Errorf(
				codes.Unauthenticated,
				"could not read auth token: %v",
				err,
			)
		}

		users, _ := queries.GetUserByToken(db.GetDBConnect(), ctx, token)
		if users.Id == 0 {
			return nil, status.Error(
				codes.PermissionDenied,
				"Please signIn or signUp",
			)
		}
		return handler(ctx, req)
	}
}
