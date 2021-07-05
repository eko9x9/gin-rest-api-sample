package middlewares

import (
	"strings"

	"github.com/eko9x9/gin-rest-api-sample/lib/utils"
	"github.com/gin-gonic/gin"
)

var jwt utils.JwtUtils

func JwtMiddleware(role *utils.JwtUserRoles) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		spHeader := strings.Split(authorization, " ")
		if spHeader[0] != "Bearer" && len(spHeader) < 1 {
			c.AbortWithStatus(401)
		}
		token := &spHeader[1]
		claim, _ := jwt.ValidateToken(*token)
		user := claim.(*utils.UserClaim)

		c.Set("user", user)

		c.Next()
	}
}
