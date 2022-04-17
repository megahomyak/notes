package main

import (
	api_views "notes/src/views/api"
	frontend_views "notes/src/views/frontend"

	"github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", frontend_views.Index)

	apiRoute := router.Group("/api")
	apiRoute.POST("/login/", api_views.Login)

    router.Run("localhost:80");
}
