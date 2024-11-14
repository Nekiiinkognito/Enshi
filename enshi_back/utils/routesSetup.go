package utils

import (
	"enshi/middleware"
	"enshi/routes"
	"enshi/routes/authRoutes"
	"enshi/routes/postsRoutes"
	"enshi/routes/userProfileRoutes"
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
	freeGroup.GET("getPost", postsRoutes.GetPost)

	// Auth group routes
	authGroup := g.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	authGroup.POST("updatePost", postsRoutes.UpdatePost)
	authGroup.POST("createPost", postsRoutes.CreatePost)
	authGroup.POST("changeUserProfile", routes.ChangeUserProfile)
	authGroup.POST("updateProfile", userProfileRoutes.UpdateUserProfile)

	authGroup.DELETE("deletePost", postsRoutes.DeletePost)

	// Admin group routes
	adminGroup := authGroup.Group("/admin/")
	adminGroup.Use(middleware.AdminMiddleware())

	adminGroup.GET("testAdmin", testAdmin)

	return nil
}
