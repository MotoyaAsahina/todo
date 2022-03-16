package main

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/MotoyaAsahina/todo/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	model.InitDB()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Debug = true

	e.Static("/", "client")

	echoAPI := e.Group("/api")
	{
		apiTasks := echoAPI.Group("/tasks")
		{
			apiTasks.GET("", router.GetTasks)
			apiTasks.POST("", router.PostTask)
			apiTasks.PUT("/:id", router.PutTask)
			apiTasks.DELETE("/:id", router.DeleteTask)
			apiTasks.PUT("/:id/done", router.PutTaskDone)
			apiTasks.PUT("/:id/undone", router.PutTaskUndone)
		}

		apiGroups := echoAPI.Group("/groups")
		{
			apiGroups.GET("", router.GetGroups)
			apiGroups.POST("", router.PostGroup)
			apiGroups.PUT("/:id", router.PutGroup)
			apiGroups.DELETE("/:id", router.DeleteGroup)
		}

		apiTags := echoAPI.Group("/tags")
		{
			apiTags.GET("", router.GetTags)
			apiTags.POST("", router.PostTag)
			apiTags.PUT("/:id", router.PutTag)
			apiTags.DELETE("/:id", router.DeleteTag)
		}
	}

	e.Logger.Fatal(e.Start(":4002"))
}
