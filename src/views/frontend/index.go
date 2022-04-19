package frontend

import (
	"net/http"
	"notes/src/config"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user, err := utils.GetUser(c, utils.WithNotes)
	if err != nil {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"user": user,
			"googleClientID": config.Config.Google.ClientID,
        })
	} else {
		templateData := gin.H{
			"user": user,
			"googleClientID": config.Config.Google.ClientID,
		}
		utils.AddCSRFToken(c, templateData)
		c.HTML(http.StatusOK, "index.tmpl", templateData)
	}
}

