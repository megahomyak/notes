package views

import (
	"net/http"
	"notes/src/config"
	"notes/src/utils"

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
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate");
	c.Header("Pragma", "no-cache");
	c.Header("Expires", "0")
	c.HTML(http.StatusOK, "index.tmpl", templateData)
}
