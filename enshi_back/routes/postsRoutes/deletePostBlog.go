package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func DeletePostBlog(c *gin.Context) {
	var queryParams db_repo.UpdatePostBlogIdParams
	postId, err := strconv.ParseInt(c.Param("post-id"), 10, 64)

	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	queryParams.BlogID = pgtype.Int8{}
	queryParams.PostID = postId

	query := db_repo.New(db_connection.Dbx)

	err = query.UpdatePostBlogId(context.Background(), queryParams)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "post has been deleted")
}
