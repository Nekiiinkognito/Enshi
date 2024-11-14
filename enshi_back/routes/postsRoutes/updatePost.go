package postsRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/checkRole"
	"enshi/middleware/getters"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UpdatePost(c *gin.Context) {
	var UpdatedPostParams db_repo.UpdatePostByPostIdParams

	if err := c.BindJSON(&UpdatedPostParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	userId, err := getters.GetUserIdFromContext(c)

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	if isOwner, _ := checkRole.IsOwnerOfThePost(
		userId,
		UpdatedPostParams.PostID,
	); !isOwner {
		rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("you are now allowed to change this"))
		return
	}

	_, err = db_repo.New(
		db_connection.Dbx,
	).UpdatePostByPostId(
		context.Background(),
		UpdatedPostParams,
	)

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "post has been updated")
}
