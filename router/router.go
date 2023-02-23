package router

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

type Handlers struct {
	Repo              model.Repository
	GoogleOAuthConfig *oauth2.Config
	WhiteList         []string
}

func (h *Handlers) SetupRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Debug = true

	e.GET("/login", h.GoogleLogin)
	e.GET("/callback", h.GoogleCallback)

	api := e.Group("/api", h.EnsureAuthorized)
	{
		apiTasks := api.Group("/tasks")
		{
			apiTasks.GET("", h.GetTasks)
			apiTasks.POST("", h.PostTask)
			apiTasks.PUT("/:id", h.PutTask)
			apiTasks.DELETE("/:id", h.DeleteTask)
			apiTasks.PUT("/:id/done", h.PutTaskDone)
			apiTasks.PUT("/:id/undone", h.PutTaskUndone)
		}

		apiGroups := api.Group("/groups")
		{
			apiGroups.GET("", h.GetGroups)
			apiGroups.POST("", h.PostGroup)
			apiGroups.PUT("/:id", h.PutGroup)
			apiGroups.DELETE("/:id", h.DeleteGroup)
			apiGroups.PUT("/:id/up", h.PutGroupUp)
			apiGroups.PUT("/:id/down", h.PutGroupDown)
			apiGroups.PUT("/:id/order", h.PutGroupOrder)
		}

		apiTags := api.Group("/tags")
		{
			apiTags.GET("", h.GetTags)
			apiTags.POST("", h.PostTag)
			apiTags.PUT("/:id", h.PutTag)
			apiTags.DELETE("/:id", h.DeleteTag)
		}
	}

	return e
}
