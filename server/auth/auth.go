package auth

import (
	"context"
	"server/model/tables"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
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

func CreateToken(user tables.AppUsers) string {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = user.Name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, _ := t.SignedString([]byte(HashPw(user.Handle + user.Password)))

	return strings.Replace(token, "_", "", -1)
}

func HashPw(pw string) string {
	// wip:ハッシュ関数
	return pw
}
