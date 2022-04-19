package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrsfMiddleware(c *gin.Context) {
	csrfTokenCookie, cookieGettingError := c.Cookie("csrf_token")
	csrfTokenFormData := c.PostForm("csrf_token")
	if cookieGettingError != nil || csrfTokenFormData == "" || csrfTokenCookie != csrfTokenFormData {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No csrf today, buddy"})
	}
}
