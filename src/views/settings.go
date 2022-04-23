package views

import (
	"net/http"
	"notes/src/models"
	"notes/src/utils"

	"github.com/gin-gonic/gin"
)

func Settings(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	c.HTML(http.StatusOK, "settings.tmpl", utils.AddCSRFToken(c, gin.H {"user": user}))
}
