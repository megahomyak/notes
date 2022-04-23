package api

import (
	"net/http"
	"notes/src/utils"

	"github.com/gin-gonic/gin"
)

func EmptyFieldError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, utils.MakeJSONError(c.Query("fieldName") + " should be filled."))
}
