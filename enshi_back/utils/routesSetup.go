package utils

import (
	"enshi/middleware"
	"enshi/routes"
	"enshi/routes/authRoutes"
	"enshi/routes/postsRoutes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func testCookie(c *gin.Context) {
	cock, _ := c.Cookie("auth_cookie")
	c.IndentedJSON(http.StatusOK, gin.H{"token": "SLESAR' U STASA " + strings.Split(cock, "_")[0]})
}

func testAdmin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "you are an admin, congrats!"})
}

func SetupRotes(g *gin.Engine) error {
	g.Use(middleware.CORSMiddleware())

	// Free group routes
	freeGroup := g.Group("/")

	freeGroup.GET("getCookie", testCookie)

	freeGroup.POST("login", authRoutes.Login)
	freeGroup.POST("registerUser", authRoutes.RegisterUser)

	// Auth group routes
	authGroup := g.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	authGroup.GET("getPost", postsRoutes.GetPost)

	authGroup.POST("createPost", postsRoutes.CreatePost)
	authGroup.POST("deletePost", postsRoutes.DeletePost)
	authGroup.POST("changeUserProfile", routes.ChangeUserProfile)

	adminGroup := authGroup.Group("/admin/")
	adminGroup.Use(middleware.AdminMiddleware())

	adminGroup.GET("testAdmin", testAdmin)

	return nil
}
