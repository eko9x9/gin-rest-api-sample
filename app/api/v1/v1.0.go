package apiv1

import (
	"github.com/eko9x9/gin-rest-api-sample/app/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")

	auth.ApplyAuthRoutes(v1)

}
