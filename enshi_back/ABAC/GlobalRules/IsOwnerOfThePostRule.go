package globalrules

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"fmt"

	"github.com/gin-gonic/gin"
)

func IsOwnerOfThePostRule(c *gin.Context) (bool, []error) {
	postId, err := getters.GetInt64Param(c, "post-id")

	if err != nil {
		return false, []error{err}
	}

	contextUserId, err := getters.GetUserIdFromContext(c)

	if err != nil {
		return false, []error{err}
	}

	post, err :=
		db_repo.New(db_connection.Dbx).
			GetPostsByPostId(context.Background(), postId)

	if err != nil {
		return false, []error{err}
	}

	if post.UserID != contextUserId {
		return false, []error{fmt.Errorf("now owner of the post")}
	}

	return true, nil
}
