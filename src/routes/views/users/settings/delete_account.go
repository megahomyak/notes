package settings

import (
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if err := models.DB.Delete(user).Error; err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusFound, "/")
}
