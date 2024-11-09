package routes

import (
	"context"
	"encoding/binary"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePost(c *gin.Context) {
	var postParams db_repo.CreatePostParams

	if err := c.BindJSON(&postParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	userId, err := getters.GetUserIdFromContext(c)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}
	postParams.UserID = userId

	if uuidForPost, err := uuid.NewV7(); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	} else {
		postParams.PostID = -int64(binary.BigEndian.Uint64(uuidForPost[8:]))
	}

	query := db_repo.New(db_connection.Dbx)
	if _, err := query.CreatePost(context.Background(), postParams); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "Post has been created!")
}
