package blogRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"enshi/utils"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	blogParams, err := utils.GetContextPayload[db_repo.CreateBlogByUserIdParams](c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	userId, err := getters.GetUserIdFromContext(c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	blogId, err := utils.GetUUIDv7AsInt64()
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	blogParams.UserID = userId
	blogParams.BlogID = blogId

	_, err = db_repo.
		New(db_connection.Dbx).
		CreateBlogByUserId(
			context.Background(),
			blogParams,
		)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "blog has been created")
}
