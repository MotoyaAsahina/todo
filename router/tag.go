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

func (h *Handlers) GetTags(c echo.Context) error {
	tags, err := h.Repo.GetTags(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, tags)
}

func (h *Handlers) PostTag(c echo.Context) error {
	req := new(PostTagRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := h.Repo.PostTag(c.Request().Context(), &model.Tag{
		ID:    uuid.New(),
		Name:  req.Name,
		Color: req.Color,
	})
	if err != nil {
		return err
	}
	return c.String(200, "")
}

func (h *Handlers) PutTag(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	req := new(PostTagRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := h.Repo.PutTag(c.Request().Context(), &model.Tag{
		ID:    id,
		Name:  req.Name,
		Color: req.Color,
	})
	if err != nil {
		return err
	}
	return c.String(200, "PutTag")
}

func (h *Handlers) DeleteTag(c echo.Context) error {
	return c.String(200, "DeleteTag")
}
