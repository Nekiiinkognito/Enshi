package utils

import (
	"context"
	"encoding/binary"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func testCookie(c *gin.Context) {
	cock, _ := c.Cookie("auth_cookie")
	c.IndentedJSON(http.StatusOK, gin.H{"token": "SLESAR' U STASA " + strings.Split(cock, "_")[0]})
}

func RegisterUser(c *gin.Context) {
	var userParams db_repo.CreateUserParams

	transaction, err := Dbx.Begin(context.Background())
	defer transaction.Rollback(context.Background())

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	if err := c.BindJSON(&userParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(userParams)
	if err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	query := db_repo.New(Dbx)
	sameNicknameOrEmailUser, _ := query.GetUserByEmailOrNickname(
		context.Background(),
		db_repo.GetUserByEmailOrNicknameParams{
			Username: userParams.Username,
			Email:    userParams.Email,
		},
	)
	if sameNicknameOrEmailUser.Username == userParams.Username {
		rest_api_stuff.ConflictAnswer(
			c,
			fmt.Errorf("username"),
		)
		return
	} else if sameNicknameOrEmailUser.Email == userParams.Email {
		rest_api_stuff.ConflictAnswer(
			c,
			fmt.Errorf("email"),
		)
		return
	}

	sqlc_transaction := Sqlc_db.WithTx(transaction)

	passwordHashSalt, err := Argon2Hasher.HashGen([]byte(userParams.Password), []byte{})
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	userParams.Password = passwordHashSalt.StringToStore

	uuid, err := uuid.NewV7()
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	userParams.UserID = int64(
		binary.BigEndian.Uint64(append(uuid[0:4], uuid[12:16]...)),
	)

	if _, err := sqlc_transaction.CreateUser(context.Background(), userParams); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	transaction.Commit(context.Background())
	rest_api_stuff.OkAnswer(c, "User has been created!")
}

func SetupRotes(g *gin.Engine) error {
	g.Use(CORSMiddleware())

	freeGroup := g.Group("/")

	// Free group routes

	freeGroup.POST("login", login)
	freeGroup.GET("getCookie", testCookie)
	freeGroup.POST("registerUser", RegisterUser)

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
