package queries

import (
	"context"
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func GetUser(engine *xorm.Engine, ctx context.Context, id uint64) (tables.AppUsers, error) {
	var user tables.AppUsers

	_, err := engine.Where(
		"id = ? and disabled = false",
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
		"id = ? and disabled = false",
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
	).And(
		"disabled = false",
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
		"handle = ? AND password = ?",
		handle,
		password,
	).And(
		"disabled = false",
	).Get(&user)

	return user, err
}

func GetMessageLogs(ctx context.Context, engine *xorm.Engine, messageLog tables.MessageLogs) (tables.MessageLogs, error) {

	_, err := engine.Where(
		"user_id = ?",
		messageLog.UserId,
	).And(
		"disabled = false",
	).Desc("id").Get(&messageLog)

	return messageLog, err
}

func GetLogToGroup(ctx context.Context, engine *xorm.Engine, logId uint64) (tables.LogToGroups, error) {
	var logToGroup tables.LogToGroups

	_, err := engine.Where(
		"log_id = ?",
		logId,
	).And(
		"disabled = false",
	).Get(&logToGroup)

	return logToGroup, err
}

func GetUserGroupByName(ctx context.Context, engine *xorm.Engine, name string) (tables.UserGroups, error) {
	var userGroups tables.UserGroups

	_, err := engine.Where(
		"name = ?",
		name,
	).And(
		"disabled = false",
	).Desc("id").Get(&userGroups)

	return userGroups, err
}

func GetUserGroupByInviteUserToGroup(ctx context.Context, engine *xorm.Engine, invite tables.InviteUserToGroups) (tables.UserGroups, error) {
	var (
		groups tables.UserGroups
	)

	engine.Alias("u").Join(
		"INNER",
		[]string{"invite_user_to_groups", "g"},
		"g.user_id = ?",
		invite.UserId,
	).Where(
		"u.id = g.group_id",
	).And(
		"g.group_id = ?",
		invite.GroupId,
	).And(
		"g.disabled = false",
	).And(
		"u.disabled = false",
	).Get(&groups)

	return groups, nil
}
