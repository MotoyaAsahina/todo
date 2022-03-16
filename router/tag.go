package router

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostTagRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func GetTags(c echo.Context) error {
	tags, err := model.GetTags(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, tags)
}

func PostTag(c echo.Context) error {
	req := new(PostTagRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := model.PostTag(c.Request().Context(), &model.Tag{
		ID:    uuid.New(),
		Name:  req.Name,
		Color: req.Color,
	})
	if err != nil {
		return err
	}
	return c.String(200, "")
}

func PutTag(c echo.Context) error {
	return c.String(200, "PutTag")
}

func DeleteTag(c echo.Context) error {
	return c.String(200, "DeleteTag")
}
