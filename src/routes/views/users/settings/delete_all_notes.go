package settings

import (
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
)

func DeleteAllNotes(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if err := models.DB.Delete(&models.Note{}, "owner_id = ?", user.ID).Error; err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusFound, "/settings/")
}
