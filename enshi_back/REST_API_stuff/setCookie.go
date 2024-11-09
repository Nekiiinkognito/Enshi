package rest_api_stuff

import (
	"github.com/gin-gonic/gin"
)

type CookieParams struct {
	Name     string
	Value    string
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

func SetCookie(c *gin.Context, params *CookieParams) {
	c.SetCookie(
		params.Name,
		params.Value,
		params.MaxAge,
		params.Path,
		params.Domain,
		params.Secure,
		params.HttpOnly,
	)
}
