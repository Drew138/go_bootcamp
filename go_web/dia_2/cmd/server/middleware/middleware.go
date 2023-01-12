package middleware

import (
	"dia_2/pkg/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var TOKEN = os.Getenv("TOKEN")

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authentication")
		if token != TOKEN {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Err(response.ErrUnauthorized))
			return
		}
		c.Next()
	}
}
