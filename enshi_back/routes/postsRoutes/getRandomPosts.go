package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRandomPost(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	postsData, err :=
		db_repo.New(db_connection.Dbx).
			GetRandomPosts(context.Background(), int32(limit))

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	result := make([]any, 0)

	for _, post := range postsData {
		result = append(result, gin.H{
			"post_id": strconv.Itoa(int(post.PostID)),
			"title":   post.Title,
			"user_id": strconv.Itoa(int(post.UserID)),
		})
	}

	c.IndentedJSON(http.StatusOK, result)

}
