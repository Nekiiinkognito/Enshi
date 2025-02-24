package getters

import (
	"enshi/auth"
	"enshi/global"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetClaimsFromContext(c *gin.Context) (auth.UserInfoJWT, error) {
	var UserInfo auth.UserInfoJWT

	claims, exists := c.Get(global.ContextTokenData)

	if !exists {
		return auth.UserInfoJWT{}, fmt.Errorf("error getting user id")
	}

	parsedUserId, err := strconv.ParseInt(
		claims.(jwt.MapClaims)["id"].(string),
		10,
		64,
	)
	if err != nil {
		return auth.UserInfoJWT{}, fmt.Errorf("error parsing user id")
	}

	UserInfo.Id = parsedUserId
	UserInfo.Username = claims.(jwt.MapClaims)["username"].(string)
	isAdmin := claims.(jwt.MapClaims)["isAdmin"].(bool)
	UserInfo.IsAdmin = isAdmin

	return UserInfo, nil

}
