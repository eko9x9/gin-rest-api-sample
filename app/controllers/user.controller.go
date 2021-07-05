package controllers

import (
	"net/http"

	"github.com/eko9x9/gin-rest-api-sample/app/models"
	"github.com/eko9x9/gin-rest-api-sample/lib/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
}

func (ctrl *UserController) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {

		userJwt := c.MustGet("user").(*utils.UserClaim)
		db := c.MustGet("db").(*gorm.DB)

		var user models.User
		if err := db.Where(models.User{Username: userJwt.Username}).First(&user).Error; err != nil {
			c.AbortWithStatus(404)

			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"name":     user.DisplayName,
			"username": user.Username,
			"join":     user.CreatedAt,
		})

	}
}
