package queries

import (
	"context"
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func DeleteInviteUserToGroupByUserId(ctx context.Context, engine *xorm.Engine, invite tables.InviteUserToGroups) (bool, error) {
	_, err := engine.Exec(
		"update invite_user_to_groups "+
			"set disabled = true "+
			"where group_id = ? AND "+
			"user_id = ?",
		invite.GroupId,
		invite.UserId,
	)
	return err != nil, err
}
