package queries

import (
	"context"
	"log"
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
	log.Println(err)

	return err == nil && affected > 0, err
}

func InsertLogToGroup(ctx context.Context, engine *xorm.Engine, logToGroup tables.LogToGroups) (bool, error) {
	affected, err := engine.Insert(&logToGroup)
	log.Println(err)

	return err == nil && affected > 0, err
}
