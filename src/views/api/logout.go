package api

import "github.com/gin-gonic/gin"

func Logout(c *gin.Context) {
	c.SetCookie("access_token", "", 0, "/", "", true, true)
}
