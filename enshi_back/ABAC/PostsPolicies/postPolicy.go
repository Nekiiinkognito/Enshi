package postspolicies

import (
	"enshi/ABAC/PostsPolicies/postRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

const (
	DELETE_POST      = "delete_post"
	UPDATE_POST      = "update_post"
	UPDATE_POST_BLOG = "update_post_blog"
	CREATE_POST      = "create_post"
	GET_POST         = "get_post"
)

func PostsPolicies(c *gin.Context) (bool, []error) {
	target, exists := c.Get("target")
	if !exists {
		return false, nil
	}

	// Permit if one permit
	switch target {
	case DELETE_POST:
		return rules.CheckRule(c, postRules.DeleteRule)

	case UPDATE_POST:
		return rules.CheckRule(c, postRules.PostUpdateRule)

	case UPDATE_POST_BLOG:
		return rules.CheckRule(c, postRules.UpdatePostBlogRule)

	case GET_POST:
		return rules.CheckRule(c, postRules.PostReadRule)

	case CREATE_POST:
		return rules.CheckRule(c, postRules.PostCreateRule)

	}

	return false, nil
}
