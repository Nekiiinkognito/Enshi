package userroutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserUsername(c *gin.Context) {
	userId, err := getters.GetInt64Param(c, "user-id")
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
	}

	userInfo, err := db_repo.New(db_connection.Dbx).GetUserUsernameById(
		context.Background(),
		userId,
	)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
	}

	c.IndentedJSON(http.StatusOK, userInfo)

}
