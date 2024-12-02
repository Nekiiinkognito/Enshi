package authRoutes

import (
	"context"
	"encoding/binary"
	rest_api_stuff "enshi/REST_API_stuff"
	"enshi/auth"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
	"enshi/global"
	"enshi/hasher"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func RegisterUser(c *gin.Context) {
	var userParams db_repo.CreateUserParams

	if err := c.BindJSON(&userParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(userParams); err != nil {
		rest_api_stuff.BadRequestAnswer(c, err)
		return
	}

	query := db_repo.New(db_connection.Dbx)
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

	transaction, err := db_connection.Dbx.Begin(context.Background())
	defer transaction.Rollback(context.Background())

	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	query_transaction := query.WithTx(transaction)

	passwordHashSalt, err := hasher.Argon2Hasher.HashGen([]byte(userParams.Password), []byte{})
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

	userParams.UserID = -int64(
		binary.BigEndian.Uint64(uuid[8:]),
	)

	if _, err := query_transaction.CreateUser(context.Background(), userParams); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	if _, err := query_transaction.CreateProfileForUser(
		context.Background(),
		userParams.UserID,
	); err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	tokenParams := auth.UserInfoJWT{
		Id:       userParams.UserID,
		Username: userParams.Username,
		IsAdmin:  false,
	}

	token, err := auth.CreateToken(tokenParams)
	if err != nil {
		rest_api_stuff.InternalErrorAnswer(c, err)
		return
	}

	cookieParams := &rest_api_stuff.CookieParams{
		Name:     "auth_cookie",
		Value:    token,
		MaxAge:   int(2 * time.Hour.Seconds()),
		Path:     global.PathForCookies,
		Domain:   global.DomainForCookies,
		Secure:   global.SecureForCookies,
		HttpOnly: global.HttpOnlyForCookies,
	}

	transaction.Commit(context.Background())
	rest_api_stuff.SetCookie(c, cookieParams)
	c.IndentedJSON(http.StatusOK, gin.H{"status": "All good", "username": userParams.Username, "id": userParams.UserID})
}
