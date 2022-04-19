package frontend

import (
	"errors"
	"net/http"
	"notes/src/models"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Note(c *gin.Context) {
	note := models.Note{}
	noteID := c.Param("note_id")
	if errors.Is(
		models.DB.Where("id = ?", noteID).Take(&note).Error,
		gorm.ErrRecordNotFound,
	) {
		c.HTML(http.StatusNotFound, "note_not_found.tmpl", gin.H{"noteID": noteID})
	} else {
		user, err := utils.GetUserByToken(c, utils.WithoutNotes)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "access_token_is_invalid.tmpl", nil)
		} else if user.ID == note.ID {
			c.HTML(http.StatusOK, "note.tmpl", gin.H{"note": note})
		} else {
			c.HTML(http.StatusForbidden, "note_is_inaccessible.tmpl", gin.H{"noteID": noteID})
		}
	}
}
