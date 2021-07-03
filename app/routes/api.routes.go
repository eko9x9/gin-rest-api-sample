package routes

import (
	"github.com/eko9x9/gin-rest-api-sample/app/routes/group"
	"github.com/gin-gonic/gin"
)

type ApiRoutes struct {
}

func (routes *ApiRoutes) ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")

	group.ApplyGroupRoutes(api)
}
