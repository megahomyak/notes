package settings

import (
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
)

func DeleteAllNotes(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	models.DB.Delete(&models.Note{}, "owner_id = ?", user.ID)
	c.Redirect(http.StatusFound, "/")
}
