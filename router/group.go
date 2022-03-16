package router

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostGroupRequest struct {
	Name string `json:"name"`
}

func GetGroups(c echo.Context) error {
	groups, err := model.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, groups)
}

func PostGroup(c echo.Context) error {
	req := new(PostGroupRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	groups, err := model.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}
	err = model.PostGroup(c.Request().Context(), &model.Group{
		Id:    uuid.New(),
		Name:  req.Name,
		Order: len(groups),
	})
	if err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func PutGroup(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	req := new(PostGroupRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := model.PutGroup(c.Request().Context(), &model.Group{
		Id:   id,
		Name: req.Name,
	})
	if err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func DeleteGroup(c echo.Context) error {
	return c.JSON(200, "DeleteGroup")
}

func PutGroupUp(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	groups, err := model.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}

	for i, group := range groups {
		if group.Id == id {
			if i == 0 {
				return c.JSON(200, nil)
			}
			err = model.PutGroupOrder(c.Request().Context(), &model.Group{
				Id:    id,
				Order: group.Order - 1,
			})
			if err != nil {
				return err
			}
			err = model.PutGroupOrder(c.Request().Context(), &model.Group{
				Id:    groups[i-1].Id,
				Order: group.Order,
			})
			if err != nil {
				return err
			}
			return c.JSON(200, nil)
		}
	}
	return c.JSON(200, nil)
}

func PutGroupDown(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	groups, err := model.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}

	for i, group := range groups {
		if group.Id == id {
			if i == len(groups)-1 {
				return c.JSON(200, nil)
			}
			err = model.PutGroupOrder(c.Request().Context(), &model.Group{
				Id:    id,
				Order: group.Order + 1,
			})
			if err != nil {
				return err
			}
			err = model.PutGroupOrder(c.Request().Context(), &model.Group{
				Id:    groups[i+1].Id,
				Order: group.Order,
			})
			if err != nil {
				return err
			}
			return c.JSON(200, nil)
		}
	}
	return c.JSON(200, nil)
}
