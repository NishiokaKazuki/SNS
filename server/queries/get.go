package queries

import (
	"context"
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func GetUser(engine *xorm.Engine, ctx context.Context, id uint64) (tables.AppUsers, error) {
	var user tables.AppUsers

	_, err := engine.Where(
		"id = ? and disbled = false",
		id,
	).Get(&user)

	return user, err
}

func GetUserByToken(engine *xorm.Engine, ctx context.Context, token string) (tables.AppUsers, error) {
	var (
		user   tables.AppUsers
		tokens tables.Tokens
	)

	_, err := engine.Where(
		"token = ?",
		token,
	).Get(&tokens)
	if err != nil {
		return user, err
	}

	_, err = engine.Where(
		"id = ? and disbled = false",
		tokens.UserId,
	).Get(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func GetUserByHandle(engine *xorm.Engine, ctx context.Context, handle string) (tables.AppUsers, error) {
	var (
		user tables.AppUsers
	)

	_, err := engine.Cols(
		"id",
	).Where(
		"handle = ?",
		handle,
	).Get(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func GetUserByPass(engine *xorm.Engine, ctx context.Context, handle string, password string) (tables.AppUsers, error) {
	var (
		user tables.AppUsers
	)

	_, err := engine.Where(
		"handle = ? And password = ?",
		handle,
		password,
	).Get(&user)
	if err != nil {
		return user, err
	}

	return user, err
}
