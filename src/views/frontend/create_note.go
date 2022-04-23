package frontend

import (
	"fmt"
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	note := &models.Note{ Name: c.PostForm("note_name"), Contents: "", OwnerID: user.ID }
	models.DB.Create(&note)
	c.Redirect(http.StatusFound, fmt.Sprintf("/note/%d", note.ID))
}
