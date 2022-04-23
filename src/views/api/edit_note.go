package api

import (
	"net/http"
	"notes/src/models"
	"notes/src/utils"

	"github.com/gin-gonic/gin"
)

func EditNote(c *gin.Context) {
	note := utils.GetNoteOr404WithJSONResponse(c)
	user := c.MustGet("user").(*models.User)
	if user.ID == note.OwnerID {
		models.DB.Model(&note).Update("contents", c.PostForm("note_contents"))
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusForbidden, utils.MakeJSONError("This note doesn't belong to you!"))
	}
}
