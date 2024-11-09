package routes

import (
	"context"
	"enshi/auth"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/global"
	"enshi/hasher"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type content struct {
		Nickname string
		Password string
	}

	var body content

	err := c.BindJSON(&body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error 1st": err.Error()})
		return
	}

	repo := db_repo.New(db_connection.Dbx)
	user, err := repo.GetUserByUsername(context.Background(), body.Nickname)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password_hash, salt, err := hasher.DecodeArgon2String(user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = hasher.Argon2Hasher.Compare(password_hash, salt, []byte(body.Password))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user_info := map[string]interface{}{
		"id":   user.UserID,
		"name": user.Username,
	}

	token, err := auth.CreateToken(user_info)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cookieName := "auth_cookie"
	cookieValue := "id=" + strconv.FormatInt(user_info["id"].(int64), 10) +
		"_nickname=" + user_info["name"].(string)
	maxAge := int(2 * time.Hour.Seconds()) // Cookie expiry time in seconds (1 hour)
	path := global.PathForCookies          // Cookie path
	domain := global.DomainForCookies      // Set domain (localhost for testing)
	secure := global.SecureForCookies      // Secure cookie (set to true in production with HTTPS)
	httpOnly := global.HttpOnlyForCookies  // HTTP only, so it can't be accessed by JavaScript

	c.Header("Authorization", token)
	c.SetCookie(cookieName, cookieValue, maxAge, path, domain, secure, httpOnly)
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}
