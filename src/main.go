package main

import (
	"notes/src/middleware"
	"notes/src/templates/loader"
	api_views "notes/src/views/api"
	frontend_views "notes/src/views/frontend"

	"github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()

	loader.LoadTemplates(router, "{{", "}}", "templates/*", "templates/template_fillers.tmpl", "templates/generic_page.tmpl")

	router.GET("/", frontend_views.Index)

	routerWithAuthenticationCheck := router.Group("/")
	routerWithAuthenticationCheck.Use(middleware.AuthMiddleware)
	routerWithAuthenticationCheck.GET("/note/:note_id", frontend_views.Note)

	apiRoute := router.Group("/api")
	apiRoute.POST("/login/", api_views.Login)

	routerWithCSRFCheck := apiRoute.Group("/")
	routerWithCSRFCheck.Use(middleware.CrsfMiddleware)
	routerWithCSRFCheck.POST("/logout/", api_views.Logout)

    router.Run("localhost:80");
}
