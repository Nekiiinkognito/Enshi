package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	var postParams struct {
		PostId int64 `json:"post_id"`
	}

	if err := c.BindJSON(&postParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	postData, err :=
		db_repo.New(db_connection.Dbx).
			GetPostsByPostId(context.Background(), postParams.PostId)

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, postData)

}
