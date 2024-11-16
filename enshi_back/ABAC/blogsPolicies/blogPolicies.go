package blogspolicies

import (
	blogrules "enshi/ABAC/blogsPolicies/blogRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

const (
	DELETE_BLOG = "delete_blog"
	UPDATE_BLOG = "update_blog"
	CREATE_BLOG = "create_blog"
	GET_BLOG    = "get_blog"
)

func BlogPolicies(c *gin.Context) (bool, []error) {
	target, exists := c.Get("target")
	if !exists {
		return false, nil
	}

	// Permit if one permit
	switch target {
	case DELETE_BLOG:
		return rules.CheckRule(c, blogrules.BlogDeleteRule)

	case UPDATE_BLOG:
		return rules.CheckRule(c, blogrules.BlogUpdateRule)

	case GET_BLOG:
		return rules.CheckRule(c, blogrules.BlogReadRule)

	case CREATE_BLOG:
		return rules.CheckRule(c, blogrules.BlogCreateRule)

	}

	return false, nil
}
