package middlewares

import (
	"fmt"
	"net/http"
	"notes/src/views/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	ResponseShouldBeJSON = true
	ResponseShouldBeHTML = false
)

func PathParametersToIntegersMiddlewareGenerator(responseShouldBeJSON bool, pathParameterNames ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		for _, pathParameterName := range pathParameterNames {
			integer, err := strconv.ParseInt(c.Param(pathParameterName), 10, 64)
			if err == nil {
				c.Set(pathParameterName, integer)
			} else {
				if responseShouldBeJSON {
					c.JSON(http.StatusBadRequest, utils.MakeJSONError(
						fmt.Sprintf("Integer from the field %s is invalid.", pathParameterName),
					))
				} else {
					c.HTML(http.StatusBadRequest, "invalid_integer.tmpl", gin.H{"pathParameterName": pathParameterName})
				}
				c.Abort()
			}
		}
	}
}
