package settings

import (
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
)

func ChangeFirstAndLastName(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if err := models.DB.Model(user).Updates(map[string]interface{}{
		"first_name": c.PostForm("first_name"),
		"last_name": c.PostForm("last_name"),
	}).Error; err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusFound, "/settings")
}
