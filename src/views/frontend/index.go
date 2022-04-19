package frontend

import (
	"net/http"
	"notes/src/config"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user := c.MustGet("user")
	templateData := gin.H{
		"user":           user,
		"googleClientID": config.Config.Google.ClientID,
	}
	if user != nil {
		utils.AddCSRFToken(c, templateData)
	}
	c.HTML(http.StatusOK, "index.tmpl", templateData)
}
