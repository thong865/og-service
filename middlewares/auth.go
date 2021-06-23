package middlewares

import "github.com/gin-gonic/gin"

func UserMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
