package queries

import (
	"context"
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func FindMessageLogsByUserId(ctx context.Context, engine *xorm.Engine, userId uint64) ([]tables.MessageLogs, error) {
	var messageLogs []tables.MessageLogs
	err := engine.Alias("m").Join(
		"INNER",
		[]string{"log_to_users", "l"},
		"l.user_id = ? AND l.is_confirmed = false",
		userId,
	).Where(
		"m.id = l.log_id",
	).And(
		"m.disabled = false",
	).Find(&messageLogs)

	return messageLogs, err
}

func FindUsersByGroupId(ctx context.Context, engine *xorm.Engine, groupId uint64) ([]tables.AppUsers, error) {
	var (
		users []tables.AppUsers
	)

	engine.Alias("u").Join(
		"INNER",
		[]string{"group_to_users", "g"},
		"g.group_id = ?",
		groupId,
	).Where(
		"u.id = g.UserId",
	).Find(&users)

	return users, nil
}
