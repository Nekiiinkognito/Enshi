package voteroutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVote(c *gin.Context) {
	var postVoteParams db_repo.GetPostVoteParams

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

	query := db_repo.New(db_connection.Dbx)
	if voteData, err := query.GetPostVote(context.Background(), postVoteParams); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"vote": voteData,
		})
	}
}
