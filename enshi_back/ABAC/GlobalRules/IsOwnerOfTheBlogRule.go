package globalrules

import (
	"context"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"fmt"

	"github.com/gin-gonic/gin"
)

func IsOwnerOfTheBlogRule(c *gin.Context) (bool, []error) {
	blogId, err := getters.GetInt64Param(c, "blog-id")

	if err != nil {
		return false, []error{err}
	}

	contextUserId, err := getters.GetUserIdFromContext(c)

	if err != nil {
		return false, []error{err}
	}

	blog, err :=
		db_repo.New(db_connection.Dbx).
			GetBlogByBlogId(context.Background(), blogId)

	if err != nil {
		return false, []error{err}
	}

	if blog.UserID != contextUserId {
		return false, []error{fmt.Errorf("now owner of the blog")}
	}

	return true, nil
}
