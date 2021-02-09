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
	result, err := engine.Exec("insert into tokens(user_id, token) "+
		" values(?, ?) "+
		"on duplicate key update "+
		"token=?",
		token.UserId,
		token.Token,
		token.Token,
	)
	affected, _ := result.RowsAffected()

	return err == nil && affected > 0, err
}

func InsertMessageLogs(ctx context.Context, engine *xorm.Engine, messageLog tables.MessageLogs) (bool, error) {
	affected, err := engine.Insert(&messageLog)

	return err == nil && affected > 0, err
}

func InsertLogToUsers(ctx context.Context, engine *xorm.Engine, logToUser tables.LogToUsers) (bool, error) {
	affected, err := engine.Insert(&logToUser)
	return err == nil && affected > 0, err
}

func InsertLogToGroup(ctx context.Context, engine *xorm.Engine, logToGroup tables.LogToGroups) (bool, error) {
	affected, err := engine.Insert(&logToGroup)
	return err == nil && affected > 0, err
}

func InsertUserGroup(ctx context.Context, engine *xorm.Engine, group tables.UserGroups) (bool, error) {
	affected, err := engine.Insert(&group)
	return err == nil && affected > 0, err
}

func InsertGroupToUser(ctx context.Context, engine *xorm.Engine, GroupToUsers tables.GroupToUsers) (bool, error) {
	affected, err := engine.Insert(&GroupToUsers)
	return err == nil && affected > 0, err
}

func InsertInviteUserToGroup(ctx context.Context, engine *xorm.Engine, invite tables.InviteUserToGroups) (bool, error) {
	affected, err := engine.Insert(&invite)
	return err == nil && affected > 0, err
}
