package utils

import (
	"enshi/middleware"
	"enshi/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func testCookie(c *gin.Context) {
	cock, _ := c.Cookie("auth_cookie")
	c.IndentedJSON(http.StatusOK, gin.H{"token": "SLESAR' U STASA " + strings.Split(cock, "_")[0]})
}

func SetupRotes(g *gin.Engine) error {
	g.Use(middleware.CORSMiddleware())

	// Free group routes
	freeGroup := g.Group("/")

	freeGroup.GET("getCookie", testCookie)

	freeGroup.POST("login", routes.Login)
	freeGroup.POST("registerUser", routes.RegisterUser)

	// Auth group routes
	authGroup := g.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	authGroup.POST("changeUserProfile", routes.ChangeUserProfile)

	return nil
}
