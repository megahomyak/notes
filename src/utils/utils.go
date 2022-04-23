package utils

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func MakeJSONError(error_text string) map[string]string {
	return map[string]string{"error": error_text}
}

func SetPermanentProtectedCookie(c *gin.Context, cookieName string, cookieContents string) {
	c.SetCookie(cookieName, cookieContents, 2147483647, "/", "", true, true)
}

func AddCSRFToken(c *gin.Context, templateData map[string]interface{}) map[string]interface{} {
	csrfToken := base64.StdEncoding.EncodeToString(MakeUniqueToken())
	SetPermanentProtectedCookie(c, "csrf_token", csrfToken)
	templateData["csrfToken"] = csrfToken
	return templateData
}

func MakeUniqueToken() []byte {
    bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
    return bytes
}
