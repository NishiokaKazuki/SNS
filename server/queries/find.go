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
	).And(
		"l.disabled = false",
	).Find(&messageLogs)

	return messageLogs, err
}

func FindAppUsersByGroupId(ctx context.Context, engine *xorm.Engine, groupId uint64) ([]tables.AppUsers, error) {
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
	).And(
		"u.disabled = false",
	).And(
		"g.disabled = false",
	).Find(&users)

	return users, nil
}

func FindUserGroupsByUserId(ctx context.Context, engine *xorm.Engine, userId uint64) ([]tables.UserGroups, error) {
	var (
		groups []tables.UserGroups
	)

	engine.Alias("u").Join(
		"INNER",
		[]string{"group_to_users", "g"},
		"g.user_id = ?",
		userId,
	).Where(
		"u.id = g.group_Id",
	).And(
		"u.disabled = false",
	).And(
		"g.disabled = false",
	).Find(&groups)

	return groups, nil
}

func FindAppUsersByInvitesGroupId(ctx context.Context, engine *xorm.Engine, groupId uint64) ([]tables.AppUsers, error) {
	var users []tables.AppUsers

	engine.Alias("u").Join(
		"INNER",
		[]string{"invite_user_to_groups", "i"},
		"i.group_id = ?",
		groupId,
	).Where(
		"u.id = i.user_id",
	).And(
		"u.disabled = false",
	).And(
		"i.disabled = false",
	).Find(&users)

	return users, nil
}

func FindInviteUserToGroupsByUserId(ctx context.Context, engine *xorm.Engine, userId uint64) ([]tables.UserGroups, error) {
	var (
		groups []tables.UserGroups
	)

	engine.Alias("u").Join(
		"INNER",
		[]string{"invite_user_to_groups", "g"},
		"g.user_id = ?",
		userId,
	).Where(
		"u.id = g.group_id",
	).And(
		"u.disabled = false",
	).And(
		"g.disabled = false",
	).Find(&groups)

	return groups, nil
}
