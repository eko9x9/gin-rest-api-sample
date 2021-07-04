package controllers

import (
	"net/http"

	"github.com/eko9x9/gin-rest-api-sample/app/form"
	"github.com/eko9x9/gin-rest-api-sample/app/models"
	"github.com/eko9x9/gin-rest-api-sample/lib/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
}

var jwt utils.JwtUtils

func (ctrl *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		var body form.LoginRequestBody

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		var user models.User
		if err := db.Where(models.User{Username: body.Username}).First(&user).Error; err != nil {
			c.AbortWithStatus(404)

			return
		}

		if !utils.HashCheck(body.Password, user.HashPassword) {
			c.AbortWithStatus(401)

			return
		}

		userInfo := utils.UserInfo{
			Name:     user.DisplayName,
			Username: user.Username,
		}

		token, err := jwt.GenerateToken(&userInfo)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully Login!",
			"token":   token,
		})

	}
}

func (ctrl *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		body := form.RegisterRequestBody{}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())

			return
		}

		hashPass, _ := utils.HashPassword(body.Password)

		newUser := &models.User{
			DisplayName:  body.DisplayName,
			Username:     body.Username,
			HashPassword: hashPass,
		}

		if dbc := db.Create(&newUser); dbc.Error != nil {
			c.JSON(http.StatusInternalServerError, dbc.Error)

			return
		}

		userInfo := utils.UserInfo{
			Name:     body.DisplayName,
			Username: body.Username,
		}
		token, err := jwt.GenerateToken(&userInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)

			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully Create Account!",
			"token":   token,
		})

	}
}
