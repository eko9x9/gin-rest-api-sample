package group

import (
	"github.com/eko9x9/gin-rest-api-sample/app/routes/group/auth"
	"github.com/gin-gonic/gin"
)

func ApplyGroupRoutes(r *gin.RouterGroup) {

	v1 := r.Group("/v1.0")

	auth.ApplyAuthRoutes(v1)
}
