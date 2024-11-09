package rest_api_stuff

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OkAnswer(c *gin.Context, message string) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": message},
	)
}

func BadRequestAnswer(c *gin.Context, err error) {
	c.IndentedJSON(
		http.StatusBadRequest,
		gin.H{"error": err.Error()},
	)
}

func InternalErrorAnswer(c *gin.Context, err error) {
	c.IndentedJSON(
		http.StatusInternalServerError,
		gin.H{"error": err.Error()},
	)
}

func ConflictAnswer(c *gin.Context, err error) {
	c.IndentedJSON(
		http.StatusConflict,
		gin.H{"error": err.Error()},
	)
}

func UnauthorizedAnswer(c *gin.Context, err error) {
	c.IndentedJSON(
		http.StatusUnauthorized,
		gin.H{"error": err.Error()},
	)
}
