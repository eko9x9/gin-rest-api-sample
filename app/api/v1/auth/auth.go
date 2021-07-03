package auth

import "github.com/gin-gonic/gin"

func ApplyAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("auth")

	auth.POST("/register", Register)
}
