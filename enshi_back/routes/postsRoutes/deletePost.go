package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("post-id"), 10, 64)

	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	query := db_repo.New(db_connection.Dbx)

	err = query.DeletePostByPostId(context.Background(), postId)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "post has been deleted")
}
