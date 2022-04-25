package notes

import (
	"fmt"
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	note := &models.Note{ Name: c.PostForm("note_name"), Contents: "", OwnerID: user.ID }
	if err := models.DB.Create(&note).Error; err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/note/%d", note.ID))
}
