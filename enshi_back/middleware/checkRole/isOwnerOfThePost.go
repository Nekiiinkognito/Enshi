package checkRole

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
)

func IsOwnerOfThePost(c *gin.Context) (bool, error) {
	postId, err := getters.GetInt64Param(c, "post-id")

	if err != nil {
		return false, err
	}

	userId, err := getters.GetUserIdFromContext(c)

	if err != nil {
		return false, err
	}

	post, err :=
		db_repo.New(db_connection.Dbx).
			GetPostsByPostId(context.Background(), postId)

	if err != nil {
		return false, err
	}

	if post.UserID != userId {
		return false, nil
	}

	return true, nil
}
