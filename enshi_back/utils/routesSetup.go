package utils

import (
	"enshi/middleware"
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

	freeGroup.POST(
		"login",
		authRoutes.Login,
	)
	freeGroup.POST(
		"users",
		authRoutes.RegisterUser,
	)
	freeGroup.GET(
		"posts/:post-id",
		postsRoutes.GetPost,
	)

	// Auth group routes
	authGroup := g.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	authGroup.PUT(
		"posts/:post-id",
		postsRoutes.UpdatePost,
	)
	authGroup.POST(
		"posts",
		postsRoutes.CreatePost,
	)
	authGroup.DELETE(
		"posts/:post-id",
		postsRoutes.DeletePost,
	)

	authGroup.PUT(
		"user-profiles",
		userProfileRoutes.UpdateUserProfile,
	)

	// Admin group routes
	adminGroup := authGroup.Group("/admin/")
	adminGroup.Use(middleware.AdminMiddleware())

	adminGroup.GET("testAdmin", testAdmin)

	return nil
}
