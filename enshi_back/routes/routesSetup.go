package routes

import (
	"enshi/middleware"
	"enshi/routes/authRoutes"
	"enshi/routes/blogRoutes"
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

	postsGroup := g.Group("/")
	postsGroup.Use(middleware.PostsMiddleware())

	postsGroup.GET(
		"posts/:post-id",
		postsRoutes.GetPost,
	)
	postsGroup.PUT(
		"posts/:post-id",
		postsRoutes.UpdatePost,
	)
	postsGroup.PUT(
		"posts/:post-id/blogs/:blog-id",
		postsRoutes.UpdatePostBlog,
	)
	postsGroup.POST(
		"posts",
		postsRoutes.CreatePost,
	)
	postsGroup.DELETE(
		"posts/:post-id",
		postsRoutes.DeletePost,
	)

	blogGroup := g.Group("/")
	blogGroup.Use(middleware.BlogsMiddleware())

	blogGroup.POST(
		"blogs",
		blogRoutes.CreateBlog,
	)

	blogGroup.PUT(
		"blogs/:blog-id",
		blogRoutes.UpdateBlog,
	)

	blogGroup.DELETE(
		"blogs/:blog-id",
		blogRoutes.DeleteBlog,
	)

	blogGroup.GET(
		"blogs/:blog-id",
		blogRoutes.GetBlog,
	)

	profilesGroup := g.Group("/")
	profilesGroup.Use(middleware.ProfileMiddleware())

	profilesGroup.PUT(
		"profiles",
		userProfileRoutes.UpdateUserProfile,
	)

	// Admin group routes
	adminGroup := g.Group("/admin/")
	adminGroup.Use(middleware.AdminMiddleware())

	adminGroup.GET("testAdmin", testAdmin)

	return nil
}
