package controllers

import (
	"net/http"

	"github.com/eko9x9/gin-rest-api-sample/app/form"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
}

func (ctrl *AuthController) Register() gin.HandlerFunc {

	return func(c *gin.Context) {
		body := form.RegisterRequestBody{}

		if err := c.ShouldBindJSON(body); err != nil {
			validate := validator.New()

			if err := validate.Struct(body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				c.Abort()

				return
			}
		}

	}
}
