package checkRole

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
)

func IsAdmin(userId int64) (bool, error) {
	user, err :=
		db_repo.New(db_connection.Dbx).
			GetUserById(context.Background(), userId)

	if err != nil || user.UserID == 0 {
		return false, err
	}

	if !user.IsAdmin {
		return false, nil
	}

	return true, nil
}
