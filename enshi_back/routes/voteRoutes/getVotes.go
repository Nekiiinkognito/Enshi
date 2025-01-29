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

func GetVotes(c *gin.Context) {
	postId, err := getters.GetInt64Param(c, "post-id")

	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	query := db_repo.New(db_connection.Dbx)
	if voteData, err := query.GetPostVotes(context.Background(), postId); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	} else {
		c.IndentedJSON(http.StatusOK, voteData)
	}
}
