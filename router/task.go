package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostTaskRequest struct {
	GroupID     uuid.UUID `json:"group_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     string    `json:"due_date"`
	Tags        []string  `json:"tags"`
}

func GetTasks(c echo.Context) error {
	return c.JSON(200, "GetTasks")
}

func PostTask(c echo.Context) error {
	var req PostTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	return c.JSON(200, "PostTasks")
}

func PutTask(c echo.Context) error {
	var req PostTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	//id := uuid.MustParse(c.Param("id"))
	return c.JSON(200, "PutTasks")
}

func DeleteTask(c echo.Context) error {
	return c.JSON(200, "DeleteTasks")
}

func PutTaskDone(c echo.Context) error {
	return c.JSON(200, "PutTaskDone")
}

func PutTaskUndone(c echo.Context) error {
	return c.JSON(200, "PutTaskUndone")
}
