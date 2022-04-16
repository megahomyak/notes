package main

import (
	api_views "notes/src/views/api"

	"github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	apiRoute := router.Group("/api")
	apiRoute.POST("/login", api_views.Login)

    router.Run("127.0.0.1:80");
}
