package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("post-id"), 10, 64)

	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	userClaims, err := getters.GetClaimsFromContext(c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	query := db_repo.New(db_connection.Dbx)
	post, err := query.GetPostsByPostId(context.Background(), postId)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	if post.UserID != userClaims.Id {
		rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("you are not the author"))
		return
	}

	// TODO: Add block of code, so admin could delete anything

	err = query.DeletePostByPostId(context.Background(), postId)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "post has been deleted")
}
