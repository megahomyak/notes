package middleware

import (
	"errors"
	"net/http"
	"notes/src/views/utils"

	"github.com/gin-gonic/gin"
)

const (
	AbortOnFailure = true
	IgnoreFailure  = false
)

func UserGetter(withNotes bool, abortOnFailure bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		user, err := utils.GetUserByToken(c, withNotes)
		if err != nil {
			if abortOnFailure {
				c.Abort()
				if errors.Is(err, &utils.AccessTokenNotFound{}) {
					c.HTML(http.StatusUnauthorized, "access_token_was_not_provided.tmpl", nil)
				} else {
					c.HTML(http.StatusUnauthorized, "access_token_is_invalid.tmpl", nil)
				}
			} else {
				c.Set("user", nil)
			}
		} else {
			c.Set("user", user)
		}
	}
}
