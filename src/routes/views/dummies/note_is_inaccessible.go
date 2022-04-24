package dummies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoteIsInaccessible(c *gin.Context) {
	c.HTML(http.StatusForbidden, "note_is_inaccessible.tmpl", gin.H{
		"noteID": c.Query("noteID"),
	})
}
