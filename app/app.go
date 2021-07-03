package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

func (app *Routes) StartServer() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
