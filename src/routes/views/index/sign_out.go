package index

import (
	"net/http"
	"notes/src/models"
	"notes/src/utils"

	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	accessTokenHash, err := utils.GetAccessTokenHash(c)
	if err != nil {
		c.HTML(http.StatusBadRequest, "access_token_was_not_provided.tmpl", nil)
	} else {
		if err := models.DB.Delete(&models.AccessToken{}, "hash = ?", accessTokenHash).Error; err != nil {
			c.Error(err)
		}
        c.SetCookie("access_token", "", 0, "/", "", true, true)
        c.Redirect(http.StatusFound, "/")
    }
}
