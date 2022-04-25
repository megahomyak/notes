package settings

import (
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	models.DB.Unscoped().Delete(user)
	c.Redirect(http.StatusFound, "/")
}
