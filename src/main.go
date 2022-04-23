package main

import (
	"notes/src/middlewares"
	views "notes/src/views"
	api "notes/src/api"
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
		getNoteRouter.GET("/note/:note_id", views.GetNote)
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
		createNoteRouter.POST("/note/", views.CreateNote)
	}

	{
		indexRouter := rootRouter.Group("/")
		indexRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithNotes, middlewares.IgnoreFailure, middlewares.ResponseShouldBeHTML,
		))
		indexRouter.GET("/", views.Index)
	}

	{
		signOutRouter := rootRouter.Group("/")
		signOutRouter.Use(middlewares.CSRFMiddleware)
		signOutRouter.POST("/sign_out/", views.SignOut)
	}

	{
		apiRouter := rootRouter.Group("/api")
		apiRouter.POST("/sign_in/", api.SignIn)
		apiRouter.GET("/empty_field_error/", api.EmptyFieldError)
		{
			editNoteRouter := apiRouter.Group("/")
			editNoteRouter.Use(middlewares.UserGetterMiddlewareGenerator(
				utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeJSON,
			))
			editNoteRouter.Use(middlewares.PathParametersToIntegersMiddlewareGenerator(
				middlewares.ResponseShouldBeJSON, "note_id",
			))
			editNoteRouter.Use(middlewares.CSRFMiddleware)
			editNoteRouter.POST("/note/:note_id/edit/", api.EditNote)
		}
	}

	{
		dummyPagesRouter := rootRouter.Group("/")
		dummyPagesRouter.GET("/empty_field_error/", views.EmptyFieldError)
		dummyPagesRouter.GET("/note_not_found/", views.NoteNotFound)
		dummyPagesRouter.GET("/note_is_inaccessible/", views.NoteIsInaccessible)
	}

	{
		deleteNoteRouter := rootRouter.Group("/")
		deleteNoteRouter.Use(middlewares.CSRFMiddleware)
		deleteNoteRouter.Use(middlewares.PathParametersToIntegersMiddlewareGenerator(
			middlewares.ResponseShouldBeHTML, "note_id",
		))
		deleteNoteRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeHTML,
		))
		deleteNoteRouter.POST("/note/:note_id/delete/", views.DeleteNote)
	}

	{
		settingsRouter := rootRouter.Group("/")
		settingsRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeHTML,
		))
		settingsRouter.GET("/settings/", views.Settings)
		{
			changeFirstAndLastNameRouter := settingsRouter.Group("/settings")
			changeFirstAndLastNameRouter.Use(middlewares.CSRFMiddleware)
			changeFirstAndLastNameRouter.POST(
				"/change_first_and_last_name/", views.ChangeFirstAndLastName,
			)
		}
	}

	// Setting up workers.

	go workers.DeleteExpiredTokensPeriodically()

	// Running the server.

	rootRouter.Run("localhost:80")
}
