package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		fmt.Println(authorization)

		c.Next()
	}
}
