package voteroutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
)

func CreateVote(c *gin.Context) {
	var postVoteParams db_repo.CreatePostVoteParams

	if err := c.BindJSON(&postVoteParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	userId, err := getters.GetUserIdFromContext(c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}
	postVoteParams.UserID = userId

	postId, err := getters.GetInt64Param(c, "post-id")
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}
	postVoteParams.PostID = postId

	query := db_repo.New(db_connection.Dbx)
	if _, err := query.CreatePostVote(context.Background(), postVoteParams); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "Vote has been created!")
}
