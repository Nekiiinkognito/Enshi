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

func GetBlog(c *gin.Context) {
	blogId, err := getters.GetInt64Param(c, "blog-id")
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	blogData, err := db_repo.New(db_connection.Dbx).
		GetBlogByBlogId(context.Background(), blogId)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, blogData)
}
