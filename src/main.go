package main

import (
	"notes/src/templates"
	api_views "notes/src/views/api"
	frontend_views "notes/src/views/frontend"

	"github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	templates.LoadTemplates(router, "{{", "}}", "templates/*", "templates/generic_page.tmpl")

	router.GET("/", frontend_views.Index)
	router.GET("/note/:note_id", frontend_views.Note)

	apiRoute := router.Group("/api")
	apiRoute.POST("/login/", api_views.Login)

    router.Run("localhost:80");
}
