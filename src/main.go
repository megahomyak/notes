package main

import (
	"notes/src/middlewares"
	"notes/src/routes/views/dummies"
	"notes/src/routes/views/index"
	"notes/src/routes/views/users/notes"
	"notes/src/routes/views/users/settings"
	"notes/src/routes/api"
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
		getNoteRouter.GET("/note/:note_id/", notes.GetNote)
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
		createNoteRouter.POST("/note/", notes.CreateNote)
	}

	{
		indexRouter := rootRouter.Group("/")
		indexRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithNotes, middlewares.IgnoreFailure, middlewares.ResponseShouldBeHTML,
		))
		indexRouter.GET("/", index.Index)
	}

	{
		signOutRouter := rootRouter.Group("/")
		signOutRouter.Use(middlewares.CSRFMiddleware)
		signOutRouter.POST("/sign_out/", index.SignOut)
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
		dummyPagesRouter.GET("/empty_field_error/", dummies.EmptyFieldError)
		dummyPagesRouter.GET("/note_not_found/", dummies.NoteNotFound)
		dummyPagesRouter.GET("/note_is_inaccessible/", dummies.NoteIsInaccessible)
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
		deleteNoteRouter.POST("/note/:note_id/delete/", notes.DeleteNote)
	}

	{
		settingsRouter := rootRouter.Group("/")
		settingsRouter.Use(middlewares.UserGetterMiddlewareGenerator(
			utils.WithoutNotes, middlewares.AbortOnFailure, middlewares.ResponseShouldBeHTML,
		))
		settingsRouter.GET("/settings/", settings.Settings)
		{
			individualSettingsRouter := settingsRouter.Group("/settings")
			individualSettingsRouter.Use(middlewares.CSRFMiddleware)
			individualSettingsRouter.POST(
				"/change_first_and_last_name/", settings.ChangeFirstAndLastName,
			)
			individualSettingsRouter.POST("/sign_out_everywhere/", settings.SignOutEverywhere)
		}
	}

	// Setting up workers.

	go workers.DeleteExpiredTokensPeriodically()

	// Running the server.

	rootRouter.Run("localhost:80")
}
