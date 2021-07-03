package app

import (
	"fmt"
	"log"
	"os"

	"github.com/eko9x9/gin-rest-api-sample/app/routes"
	"github.com/eko9x9/gin-rest-api-sample/database"
	"github.com/gin-gonic/gin"
)

type App struct {
}

func (app *App) StartServer() {
	r := gin.Default()

	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	} else {
		fmt.Println("Successfully connect database")
	}

	r.Use(database.Inject(db))
	r.Use(gin.Logger())

	apiRoutes := new(routes.ApiRoutes)
	apiRoutes.ApplyRoutes(r)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
