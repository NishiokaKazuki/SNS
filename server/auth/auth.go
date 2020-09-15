package auth

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// AuthorizationKey は認証トークンに対応するキーを表す
	AuthorizationKey = "authorization"
)

func Auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not read auth token: %v",
			err,
		)
	}

	// デモ実装であるため、署名は検証しない。
	parser := new(jwt.Parser)
	parsedToken, _, err := parser.ParseUnverified(token, &jwt.StandardClaims{})
	if err != nil {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not parsed auth token: %v",
			err,
		)
	}

	return setToken(ctx, parsedToken.Claims.(*jwt.StandardClaims)), nil
}

func setToken(ctx context.Context, token *jwt.StandardClaims) context.Context {
	return context.WithValue(ctx, AuthorizationKey, token)
}

func GetToken(ctx context.Context) *jwt.StandardClaims {
	return ctx.Value(AuthorizationKey).(*jwt.StandardClaims)
}
