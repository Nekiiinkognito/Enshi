package checkRole

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) (bool, error) {
	userId, err := getters.GetUserIdFromContext(c)

	if err != nil {
		return false, err
	}

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
