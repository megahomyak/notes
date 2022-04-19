package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrsfMiddleware(c *gin.Context) {
	csrfTokenCookie, cookieGettingError := c.Cookie("csrf_token")
	fmt.Print(csrfTokenCookie, "\n")
	csrfTokenFormData := c.PostForm("csrf_token")
	fmt.Print(csrfTokenFormData, "\n")
	if cookieGettingError != nil || csrfTokenFormData == "" || csrfTokenCookie != csrfTokenFormData {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No csrf today, buddy"})
	}
}
