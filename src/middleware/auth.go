package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "access_token_was_not_provided.tmpl", nil)
		c.Abort()
	} else {
		c.Set("accessToken", accessToken)
	}
}
