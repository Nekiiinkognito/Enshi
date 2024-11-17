package getters

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetContextPayload[T any](c *gin.Context) (T, error) {
	var params T

	if err := c.BindJSON(&params); err != nil {
		return params, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(params); err != nil {
		return params, err
	}

	return params, nil
}
