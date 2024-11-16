package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
)

func UpdatePost(c *gin.Context) {
	var UpdatedPostParams db_repo.UpdatePostByPostIdParams

	if err := c.BindJSON(&UpdatedPostParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	postId, err := getters.GetInt64Param(c, "post-id")

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	UpdatedPostParams.PostID = postId

	_, err = db_repo.New(
		db_connection.Dbx,
	).UpdatePostByPostId(
		context.Background(),
		UpdatedPostParams,
	)

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "post has been updated")
}
