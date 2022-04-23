package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFormFieldsValidatorMiddlewareGenerator(responseShouldBeJSON bool, fieldNames ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		for _, fieldName := range fieldNames {
			if c.PostForm(fieldName) == "" {
				if responseShouldBeJSON {
					c.Redirect(http.StatusFound, "/api/empty_field_error/?fieldName=" + fieldName)
				} else {
					c.Redirect(http.StatusFound, "/empty_field_error/?fieldName=" + fieldName)
				}
				c.Abort()
				return
			}
		}
	}
}
