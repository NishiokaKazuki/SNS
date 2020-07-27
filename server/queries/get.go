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
