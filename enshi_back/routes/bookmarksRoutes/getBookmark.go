package bookmarksroutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetBookmark(c *gin.Context) {
	var bookmarkParams db_repo.GetBookmarkTimestampParams

	if err := c.BindJSON(&bookmarkParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	userId, err := getters.GetUserIdFromContext(c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}
	bookmarkParams.UserID = userId

	query := db_repo.New(db_connection.Dbx)
	if timestamp, err := query.GetBookmarkTimestamp(context.Background(), bookmarkParams); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	} else {
		if timestamp.Valid {
			c.IndentedJSON(http.StatusOK, gin.H{
				"isBookmarked": timestamp.Valid,
				"bookmarkedAt": timestamp.Time,
			})
			return
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"isBookmarked": timestamp.Valid,
				"bookmarkedAt": time.Unix(1<<63-1, 0).UTC(),
			})
		}
	}
}
