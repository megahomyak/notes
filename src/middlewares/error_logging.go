package middlewares

import (
	"notes/src/logging"

	"github.com/gin-gonic/gin"
)

func ErrorLogger(c *gin.Context) {
	c.Next()
	for _, errorText := range c.Errors.Errors() {
		logging.Log(errorText)
	}
}
