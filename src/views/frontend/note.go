package frontend

import (
	"net/http"
	"notes/src/models"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
)

func Note(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	note := utils.GetNoteOr404WithHTMLResponse(c)
	if note == nil {
		return
	}
	if user.ID == note.ID {
		c.HTML(http.StatusOK, "note.tmpl", utils.AddCSRFToken(c, gin.H{"note": note}))
	} else {
		c.HTML(http.StatusForbidden, "note_is_inaccessible.tmpl", gin.H{"noteID": note.ID})
	}
}
