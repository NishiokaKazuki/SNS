package queries

import (
	"context"
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func UpdatelogToUsers(ctx context.Context, engine *xorm.Engine, logToUser tables.LogToUsers, logIds []uint64, userId uint64) (bool, error) {
	affected, err := engine.Cols(
		"is_confirmed",
	).In(
		"log_id",
		logIds,
	).Where(
		"user_id = ?",
		userId,
	).Update(&logToUser)

	return err == nil && affected > 0, err
}
