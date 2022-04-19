package main

import (
	"notes/src/middleware"
	"notes/src/templates/loader"
	api_views "notes/src/views/api"
	frontend_views "notes/src/views/frontend"
	"notes/src/views/utils"
	"notes/src/workers"

	"github.com/gin-gonic/gin"
)


func main() {
	// Setting up routes.

    router := gin.Default()

	loader.LoadTemplates(router, "{{", "}}", "templates/*", "templates/template_fillers.tmpl", "templates/generic_page.tmpl")

	{
		routerWithRequiredUser := router.Group("/")
		routerWithRequiredUser.Use(middleware.UserGetter(utils.WithoutNotes, middleware.AbortOnFailure))
		routerWithRequiredUser.GET("/note/:note_id", frontend_views.Note)
	}

	{
		routerWithOptionalUser := router.Group("/")
		routerWithOptionalUser.Use(middleware.UserGetter(utils.WithNotes, middleware.IgnoreFailure))
		routerWithOptionalUser.GET("/", frontend_views.Index)
	}

	{
		apiRoute := router.Group("/api")
		apiRoute.POST("/sign_in/", api_views.SignIn)

		{
			routerWithCSRFCheck := apiRoute.Group("/")
			routerWithCSRFCheck.Use(middleware.CrsfMiddleware)
			routerWithCSRFCheck.POST("/sign_out/", api_views.SignOut)
		}
	}

	// Setting up workers.

	go workers.DeleteExpiredTokensPeriodically()

	// Running the server.

    router.Run("localhost:80");
}
