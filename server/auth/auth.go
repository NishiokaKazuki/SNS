package auth

import (
	"context"
	"server/model/db"
	"server/model/tables"
	"server/queries"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Auth(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func StreamServerAuthorized(ctx context.Context) (context.Context, error) {
	token, err := GetToken(ctx)
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
	return ctx, nil
}

func GetToken(ctx context.Context) (string, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	return token, err
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
