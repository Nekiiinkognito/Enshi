package postRules

import (
	"github.com/gin-gonic/gin"
)

// Only owner of the post can change it
func PostReadRule(c *gin.Context) (bool, error) {
	return true, nil
}
