package frontend

import (
	"errors"
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Note(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	note := models.Note{}
	noteID := c.Param("note_id")
	if errors.Is(
		models.DB.Where("id = ?", noteID).Take(&note).Error,
		gorm.ErrRecordNotFound,
	) {
		c.HTML(http.StatusNotFound, "note_not_found.tmpl", gin.H{"noteID": noteID})
	} else {
		if user.ID == note.ID {
			c.HTML(http.StatusOK, "note.tmpl", gin.H{"note": note})
		} else {
			c.HTML(http.StatusForbidden, "note_is_inaccessible.tmpl", gin.H{"noteID": noteID})
		}
	}
}
