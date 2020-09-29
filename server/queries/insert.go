package queries

import (
	"context"
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func InsertAppUser(engine *xorm.Engine, ctx context.Context, appUser tables.AppUsers) (bool, error) {

	affected, err := engine.Cols(
		"handle",
		"password",
		"name",
	).Insert(&appUser)

	return err == nil && affected > 0, err
}

func InsertTokens(engine *xorm.Engine, ctx context.Context, token tables.Tokens) (bool, error) {
	affected, err := engine.Cols(
		"user_id",
		"tokens",
	).Insert(&token)

	return err == nil && affected > 0, err
}
