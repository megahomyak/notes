package frontend

import (
	"errors"
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Note(c *gin.Context) {
	note := models.Note{}
	noteID := c.Param("note_id")
	if errors.Is(
		models.DB.Where("id = ?", noteID).Preload("Owner").Take(&note).Error,
		gorm.ErrRecordNotFound,
	) {
		c.HTML(http.StatusNotFound, "note_not_found.tmpl", gin.H{"noteID": noteID})
	} else {
		if note.Owner.AccessToken.Valid {
			accessToken, err := c.Cookie("access_token")
			if err != nil {
				c.HTML(http.StatusBadRequest, "token_was_not_provided.html", nil)
				return
			}
			if note.Owner.AccessToken.String == accessToken {
				c.HTML(http.StatusOK, "note.tmpl", gin.H{"note": note})
			} else {
				c.HTML(http.StatusForbidden, "note_is_inaccessible.tmpl", gin.H{"noteID": noteID})
			}
		}
	}
}
