package router

import (
	mid "github.com/MotoyaAsahina/todo/middleware"
	"github.com/MotoyaAsahina/todo/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	Repo model.Repository
}

func (h *Handlers) SetupRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Debug = true

	router.SetupGoogleOauth2()
	router.ResetNotifications()

	e.GET("/login", router.GoogleLogin)
	e.GET("/callback", router.GoogleCallback)

	api := e.Group("/api", mid.EnsureAuthorized)
	{
		apiTasks := api.Group("/tasks")
		{
			apiTasks.GET("", router.GetTasks)
			apiTasks.POST("", router.PostTask)
			apiTasks.PUT("/:id", router.PutTask)
			apiTasks.DELETE("/:id", router.DeleteTask)
			apiTasks.PUT("/:id/done", router.PutTaskDone)
			apiTasks.PUT("/:id/undone", router.PutTaskUndone)
		}

		apiGroups := api.Group("/groups")
		{
			apiGroups.GET("", router.GetGroups)
			apiGroups.POST("", router.PostGroup)
			apiGroups.PUT("/:id", router.PutGroup)
			apiGroups.DELETE("/:id", router.DeleteGroup)
			apiGroups.PUT("/:id/up", router.PutGroupUp)
			apiGroups.PUT("/:id/down", router.PutGroupDown)
			apiGroups.PUT("/:id/order", router.PutGroupOrder)
		}

		apiTags := api.Group("/tags")
		{
			apiTags.GET("", router.GetTags)
			apiTags.POST("", router.PostTag)
			apiTags.PUT("/:id", router.PutTag)
			apiTags.DELETE("/:id", router.DeleteTag)
		}
	}
}
