package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	postId, err := getters.GetInt64Param(c, "post-id")

	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	postData, err :=
		db_repo.New(db_connection.Dbx).
			GetPostsByPostId(context.Background(), postId)

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, postData)

}
