package globalrules

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"fmt"

	"github.com/gin-gonic/gin"
)

func IsAdminRule(c *gin.Context) (bool, []error) {
	contextUserId, err := getters.GetUserIdFromContext(c)

	if err != nil {
		return false, []error{err}
	}

	user, err :=
		db_repo.New(db_connection.Dbx).
			GetUserById(context.Background(), contextUserId)

	if err != nil || user.UserID == 0 {
		return false, []error{err}
	}

	if !user.IsAdmin {
		return false, []error{fmt.Errorf("not admin")}
	}

	return true, nil
}
