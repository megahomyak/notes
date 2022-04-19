package api

import (
	"notes/src/models"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	accessTokenHash, err := utils.GetAccessTokenHash(c)
	if err != nil {
		utils.MakeJSONError("access_token wasn't provided!")
	} else {
		models.DB.Delete(&models.AccessToken{}, "hash = ?", accessTokenHash)
		c.SetCookie("access_token", "", 0, "/", "", true, true)
	}
}
