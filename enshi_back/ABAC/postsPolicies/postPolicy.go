package postspolicies

import (
	postRules "enshi/ABAC/postsPolicies/postRules"

	"github.com/gin-gonic/gin"
)

const (
	DELETE_POST = "delete_post"
	UPDATE_POST = "update_post"
	CREATE_POST = "create_post"
	GET_POST    = "get_post"
)

func checkRule(
	c *gin.Context,
	ruleChecker func(*gin.Context) (bool, error),
) (bool, error) {
	IsAllowed, err := ruleChecker(c)
	if err != nil {
		return false, err
	}

	return IsAllowed, nil
}

func PostsPolicies(c *gin.Context) (bool, error) {
	target, exists := c.Get("target")
	if !exists {
		return false, nil
	}

	switch target {
	case DELETE_POST:
		return checkRule(c, postRules.DeleteRule)

	case UPDATE_POST:
		return checkRule(c, postRules.PostUpdateRule)

	case GET_POST:
		return checkRule(c, postRules.PostReadRule)

	case CREATE_POST:
		return checkRule(c, postRules.PostCreateRule)

	}

	return false, nil
}
