package auth

import (
	"github.com/eko9x9/gin-rest-api-sample/app/controllers"
	"github.com/gin-gonic/gin"
)

func ApplyAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("auth")

	authController := new(controllers.AuthController)

	auth.POST("/register", authController.Register())
}
