package main

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/MotoyaAsahina/todo/router"
)

func main() {
	gormRepo := &model.GormRepository{}
	err := gormRepo.InitDB()
	if err != nil {
		panic(err)
	}

	handler := &router.Handlers{
		Repo: gormRepo,
	}
	handler.SetupGoogleOauth2()
	handler.ResetNotifications()

	e := handler.SetupRoutes()
	e.Logger.Fatal(e.Start(":8010"))
}
