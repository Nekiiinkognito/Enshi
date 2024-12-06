package blogRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserBlogs(c *gin.Context) {
	userId, err := getters.GetUserIdFromContext(c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	blogData, err := db_repo.New(db_connection.Dbx).
		GetBlogsByUserId(context.Background(), userId)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, blogData)
}
