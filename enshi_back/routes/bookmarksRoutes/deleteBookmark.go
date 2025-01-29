package bookmarksroutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
)

func DeleteBookmark(c *gin.Context) {
	var bookmarkParams db_repo.DeleteBookmarkParams

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
	if err := query.DeleteBookmark(context.Background(), bookmarkParams); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "Bookmark has been deleted!")
}
