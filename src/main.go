package main

import (
	"notes/src/middlewares"
	"notes/src/templates/loader"
	api_views "notes/src/views/api"
	frontend_views "notes/src/views/frontend"
	"notes/src/views/utils"
	"notes/src/workers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setting up routes.

	rootRouter := gin.Default()

	loader.LoadTemplates(
		rootRouter, "{{", "}}",
		"templates/*", "templates/template_fillers.tmpl", "templates/generic_page.tmpl",
	)

	{
		noteRouter := rootRouter.Group("/")
		noteRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeHTML,
		))
		noteRouter.Use(middlewares.PathParametersToIntegersMiddlewareGenerator(
			middlewares.ResponseShouldBeHTML, "note_id",
		))
		noteRouter.GET("/note/:note_id", frontend_views.Note)
	}

	{
		indexRouter := rootRouter.Group("/")
		indexRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithNotes, middlewares.IgnoreFailure, middlewares.ResponseShouldBeHTML,
		))
		indexRouter.GET("/", frontend_views.Index)
	}

	{
		apiRouter := rootRouter.Group("/api")
		apiRouter.POST("/sign_in/", api_views.SignIn)
		{
			routerWithCSRFCheck := apiRouter.Group("/")
			routerWithCSRFCheck.Use(middlewares.CSRFMiddleware)
			routerWithCSRFCheck.POST("/sign_out/", api_views.SignOut)
			{
				editNoteRouter := routerWithCSRFCheck.Group("/")
				editNoteRouter.Use(middlewares.UserGetterMiddlewareGenerator(
					utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeJSON,
				))
				editNoteRouter.Use(middlewares.PathParametersToIntegersMiddlewareGenerator(
					middlewares.ResponseShouldBeJSON, "note_id",
				))
				editNoteRouter.POST("/note/:note_id/edit/", api_views.EditNote)
			}
		}
	}

	// Setting up workers.

	go workers.DeleteExpiredTokensPeriodically()

	// Running the server.

	rootRouter.Run("localhost:80")
}
