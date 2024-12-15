package routes

import (
	"enshi/middleware"
	"enshi/middleware/getters"
	"enshi/routes/authRoutes"
	"enshi/routes/blogRoutes"
	bookmarksroutes "enshi/routes/bookmarksRoutes"
	"enshi/routes/postsRoutes"
	"enshi/routes/userProfileRoutes"
	userroutes "enshi/routes/userRoutes"
	voteroutes "enshi/routes/voteRoutes"
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

func testAuth(c *gin.Context) {
	userInfo, err := getters.GetClaimsFromContext(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "you are not logged in"})

	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message":  "you are logged in, congrats!",
			"username": userInfo.Username,
			"is_admin": userInfo.IsAdmin,
			"id":       userInfo.Id,
		},
	)
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

	postsGroup.GET(
		"posts/random",
		postsRoutes.GetRandomPost,
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
	postsGroup.DELETE(
		"posts/:post-id/blogs",
		postsRoutes.DeletePostBlog,
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

	bookmarksGroup := g.Group("/")
	bookmarksGroup.Use(middleware.BookmarksMiddleware())

	bookmarksGroup.POST(
		"bookmarks/:post-id",
		bookmarksroutes.CreateBookmark,
	)

	bookmarksGroup.DELETE(
		"bookmarks/:post-id",
		bookmarksroutes.DeleteBookmark,
	)

	bookmarksGroup.GET(
		"bookmarks/:post-id",
		bookmarksroutes.GetBookmark,
	)

	postVoteGroup := g.Group("/")
	postVoteGroup.Use(middleware.PostVotesMiddleware())

	postVoteGroup.POST(
		"post-votes/:post-id",
		voteroutes.CreateVote,
	)

	postVoteGroup.DELETE(
		"post-votes/:post-id",
		voteroutes.DeleteVote,
	)

	postVoteGroup.GET(
		"post-vote/:post-id",
		voteroutes.GetVote,
	)

	postVoteGroup.GET(
		"post-votes/:post-id",
		voteroutes.GetVotes,
	)

	// Admin group routes
	adminGroup := g.Group("/admin/")
	adminGroup.Use(middleware.AdminMiddleware())

	adminGroup.GET("check", testAdmin)

	authGroup := g.Group("/auth/")
	authGroup.Use(middleware.AuthMiddleware())
	authGroup.GET("check", testAuth)

	temporal := g.Group("/")
	temporal.Use(middleware.AuthMiddleware())

	temporal.GET(
		"/user/blogs",
		blogRoutes.GetUserBlogs,
	)

	freeGroup.GET(
		"/user/:user-id",
		userroutes.GetUserUsername,
	)

	return nil
}
