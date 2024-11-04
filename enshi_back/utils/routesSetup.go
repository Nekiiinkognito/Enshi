package utils

import (
	"context"
	db_repo "enshi/db/go_queries"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func testCookie(c *gin.Context) {
	cock, _ := c.Cookie("auth_cookie")
	c.IndentedJSON(http.StatusOK, gin.H{"token": "SLESAR' U STASA " + strings.Split(cock, "_")[0]})
}

func SetupRotes(g *gin.Engine) error {
	g.Use(CORSMiddleware())

	freeGroup := g.Group("/")

	// Free group routes

	freeGroup.POST("login", login)
	freeGroup.GET("getCookie", testCookie)

	authGroup := g.Group("/")
	authGroup.Use(AuthMiddleware())

	// Auth group routes

	return nil
}

func login(c *gin.Context) {
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

	repo := db_repo.New(Dbx)
	user, err := repo.GetUserByUsername(context.Background(), body.Nickname)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password_hash, salt, err := DecodeArgon2String(user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Argon2Hasher.Compare(password_hash, salt, []byte(body.Password))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user_info := map[string]interface{}{
		"id":   user.UserID,
		"name": user.Username,
	}

	token, err := CreateToken(user_info)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cookieName := "auth_cookie"
	cookieValue := "id=" + strconv.FormatInt(user_info["id"].(int64), 10) +
		"_nickname=" + user_info["name"].(string)
	maxAge := int(2 * time.Hour.Seconds()) // Cookie expiry time in seconds (1 hour)
	path := "/"                            // Cookie path
	domain := "localhost"                  // Set domain (localhost for testing)
	secure := false                        // Secure cookie (set to true in production with HTTPS)
	httpOnly := false                      // HTTP only, so it can't be accessed by JavaScript
	// LookupEnv(&domain, "DOMAIN")

	c.Header("Authorization", token)
	c.SetCookie(cookieName, cookieValue, maxAge, path, domain, secure, httpOnly)
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, Authorization, accept, origin, Cache-Control, X-Requested-With, Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Token, Uid, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		claims, err := ValidateToken(token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
			c.Abort()
			return
		}

		// Claims -> data stored in token
		c.Set("id", claims["id"])
		c.Set("claims", claims)
		c.Next()

	}
}
