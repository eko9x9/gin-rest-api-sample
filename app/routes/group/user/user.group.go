package user

import (
	"github.com/eko9x9/gin-rest-api-sample/app/controllers"
	"github.com/eko9x9/gin-rest-api-sample/lib/middlewares"
	"github.com/eko9x9/gin-rest-api-sample/lib/utils"
	"github.com/gin-gonic/gin"
)

func ApplyUserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	role := utils.None
	userGroup.Use(middlewares.JwtMiddleware(&role))

	userCtrl := new(controllers.UserController)

	userGroup.GET("/profile", userCtrl.GetProfile())

}
