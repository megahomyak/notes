package main

import (
	"notes/src/middlewares"
	api_views "notes/src/views/api"
	frontend_views "notes/src/views/frontend"
	"notes/src/utils"
	"notes/src/workers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setting up routes.

	rootRouter := gin.Default()

	utils.LoadTemplates(
		rootRouter, "{{", "}}",
		"templates/*", "templates/template_fillers.tmpl", "templates/generic_page.tmpl",
	)

	{
		getNoteRouter := rootRouter.Group("/")
		getNoteRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeHTML,
		))
		getNoteRouter.Use(middlewares.PathParametersToIntegersMiddlewareGenerator(
			middlewares.ResponseShouldBeHTML, "note_id",
		))
		getNoteRouter.GET("/note/:note_id", frontend_views.GetNote)
	}

	{
		createNoteRouter := rootRouter.Group("/")
		createNoteRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeHTML,
		))
		createNoteRouter.Use(middlewares.CSRFMiddleware)
		createNoteRouter.Use(middlewares.PostFormFieldsValidatorMiddlewareGenerator(
			middlewares.ResponseShouldBeHTML, "note_name",
		))
		createNoteRouter.POST("/note/", frontend_views.CreateNote)
	}

	{
		indexRouter := rootRouter.Group("/")
		indexRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithNotes, middlewares.IgnoreFailure, middlewares.ResponseShouldBeHTML,
		))
		indexRouter.GET("/", frontend_views.Index)
	}

	{
		signOutRouter := rootRouter.Group("/")
		signOutRouter.Use(middlewares.CSRFMiddleware)
		signOutRouter.POST("/sign_out/", frontend_views.SignOut)
	}

	{
		apiRouter := rootRouter.Group("/api")
		apiRouter.POST("/sign_in/", api_views.SignIn)
		apiRouter.GET("/empty_field_error/", api_views.EmptyFieldError)
		{
			editNoteRouter := apiRouter.Group("/")
			editNoteRouter.Use(middlewares.UserGetterMiddlewareGenerator(
				utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeJSON,
			))
			editNoteRouter.Use(middlewares.PathParametersToIntegersMiddlewareGenerator(
				middlewares.ResponseShouldBeJSON, "note_id",
			))
			editNoteRouter.Use(middlewares.CSRFMiddleware)
			editNoteRouter.POST("/note/:note_id/edit/", api_views.EditNote)
		}
	}

	{
		rootRouter.GET("/empty_field_error/", frontend_views.EmptyFieldError)
	}

	// Setting up workers.

	go workers.DeleteExpiredTokensPeriodically()

	// Running the server.

	rootRouter.Run("localhost:80")
}
