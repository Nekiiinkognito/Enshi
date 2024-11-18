package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func UpdatePostBlog(c *gin.Context) {
	var UpdatedPostParams db_repo.UpdatePostBlogIdParams

	postId, err := getters.GetInt64Param(c, "post-id")

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	blogId, err := getters.GetInt64Param(c, "blog-id")

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	UpdatedPostParams.PostID = postId
	UpdatedPostParams.BlogID = pgtype.Int8{
		Valid: true,
		Int64: blogId,
	}

	transaction, err := db_connection.Dbx.Begin(context.Background())
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}
	defer transaction.Rollback(context.Background())

	err = db_repo.New(
		transaction,
	).UpdatePostBlogId(
		context.Background(),
		UpdatedPostParams,
	)

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	transaction.Commit(context.Background())
	rest_api_stuff.OkAnswer(c, "post has been updated")
}
