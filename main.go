package main

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/MotoyaAsahina/todo/router"
)

func main() {
	model.InitDB()

	gormRepo := model.Repository()

	handler := &router.Handlers{
		Repo: model.NewRepository(),
	}

	e.Logger.Fatal(e.Start(":8010"))
}
