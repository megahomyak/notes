package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmptyFieldError(c *gin.Context) {
	c.HTML(http.StatusBadRequest, "empty_field_error.tmpl", gin.H{
		"fieldName": c.Query("fieldName"),
	})
}
