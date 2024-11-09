package routes

import (
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"

	"github.com/gin-gonic/gin"
)

func ChangeUserProfile(c *gin.Context) {
	var userProfileParams db_repo.UpdateProfileByUserIdParams

	if err := c.BindJSON(&userProfileParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
	}

}
