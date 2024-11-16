package userProfileRoutes

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/middleware/getters"
	"enshi/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUserProfile(c *gin.Context) {
	newProfile, err :=
		utils.GetContextPayload[db_repo.UpdateProfileByUserIdParams](c)

	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	userId, err := getters.GetUserIdFromContext(c)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	newProfile.UserID = userId

	if _, err := db_repo.New(db_connection.Dbx).UpdateProfileByUserId(
		context.Background(),
		newProfile,
	); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	rest_api_stuff.OkAnswer(c, "profile was updated")
}
