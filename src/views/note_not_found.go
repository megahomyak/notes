package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "note_not_found.tmpl", gin.H{
		"noteID": c.Query("noteID"),
	})
}
