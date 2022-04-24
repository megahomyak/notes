package notes

import (
	"errors"
	"net/http"
	"notes/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteNote(c *gin.Context) {
	noteID := c.MustGet("note_id").(int64)
	note := &models.Note{}
	if errors.Is(models.DB.Where("id = ?", noteID).Take(&note).Error, gorm.ErrRecordNotFound) {
		// I know that the line below looks weird, but just so you understand it better:
		// I have to redirect to some GET-endpoint to prevent a chance of the POST-form being
		// re-sent, and the only HTTP status code that tells the browser to redirect to a GET
		// endpoint AND is semantically non-permanent is FOUND, so I have no choice other than
		// using HTTP FOUND to redirect to a NOT FOUND page. Thank you for your attention.
		c.Redirect(http.StatusFound, "/note_not_found/?noteID=" + strconv.Itoa(int(noteID)))
	} else {
		user := c.MustGet("user").(*models.User)
		if user.ID == note.OwnerID {
			models.DB.Delete(&note)
			c.Redirect(http.StatusFound, "/")
		} else {
			// Ugly. Ugly, but I decided to make this site as static as possible. I swear, it is
			// just for the sake of experiment, usually I don't do such things, check
			// `django_to_do_list` if you don't believe me.
			c.Redirect(http.StatusFound, "/note_is_inaccessible/?noteID=" + strconv.Itoa(int(noteID)))
		}
	}
}
