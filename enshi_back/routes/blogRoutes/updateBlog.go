package blogRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
)

func UpdateBlog(c *gin.Context) {
	newBlogParams, err :=
		getters.GetContextPayload[db_repo.UpdateBlogInfoByBlogIdParams](c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	blogId, err := getters.GetInt64Param(c, "blog-id")
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	newBlogParams.BlogID = blogId

	transaction, err := db_connection.Dbx.Begin(context.Background())
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}
	defer transaction.Rollback(context.Background())

	_, err = db_repo.New(transaction).
		UpdateBlogInfoByBlogId(context.Background(), newBlogParams)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}
	transaction.Commit(context.Background())

	rest_api_stuff.OkAnswer(c, "blog has been updated")

}
