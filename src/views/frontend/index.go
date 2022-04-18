package frontend

import (
	"net/http"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user, err := utils.GetUser(c, utils.WithNotes)
	if err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"userIsLoggedIn": false,
		})
	} else {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "userIsLoggedIn": true,
            "notes": user.Notes,
        })
	}
}

