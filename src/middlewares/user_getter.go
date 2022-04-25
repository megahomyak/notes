package middlewares

import (
	"errors"
	"net/http"
	"notes/src/utils"

	"github.com/gin-gonic/gin"
)

const (
	AbortOnFailure = true
	IgnoreFailure  = false
)

func UserGetterMiddlewareGenerator(abortOnFailure bool, responseShouldBeJSON bool) func(*gin.Context) {
	return func(c *gin.Context) {
		user, err := utils.GetUserByToken(c)
		if err != nil {
			if abortOnFailure {
				if responseShouldBeJSON {
					if errors.Is(err, &utils.AccessTokenNotFound{}) {
						c.HTML(http.StatusUnauthorized, "access_token_was_not_provided.tmpl", nil)
					} else {
						c.HTML(http.StatusUnauthorized, "access_token_is_invalid.tmpl", nil)
					}
				} else {
					if errors.Is(err, &utils.AccessTokenNotFound{}) {
						c.JSON(http.StatusUnauthorized, utils.MakeJSONError("access_token was not provided."))
					} else {
						c.JSON(http.StatusUnauthorized, utils.MakeJSONError("access_token is invalid."))
					}
				}
				c.Abort()
			} else {
				c.Set("user", nil)
			}
		} else {
			c.Set("user", user)
		}
	}
}
