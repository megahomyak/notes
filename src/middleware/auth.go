package middleware

import (
	"net/http"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			utils.MakeJSONError("access_token is not provided!"),
		)
	} else {
		c.Set("accessToken", accessToken)
	}
}
