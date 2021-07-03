package app

import (
	"os"

	"github.com/eko9x9/gin-rest-api-sample/app/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
}

func (app *App) StartServer() {
	r := gin.Default()
	r.Use(gin.Logger())

	apiRoutes := new(routes.ApiRoutes)
	apiRoutes.ApplyRoutes(r)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
