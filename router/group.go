package router

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostGroupRequest struct {
	Name string `json:"name"`
}

type PutGroupOrderRequest struct {
	Order int `json:"order"`
}

func (h *Handlers) GetGroups(c echo.Context) error {
	groups, err := h.Repo.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, groups)
}

func (h *Handlers) PostGroup(c echo.Context) error {
	req := new(PostGroupRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	groups, err := h.Repo.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}
	err = h.Repo.PostGroup(c.Request().Context(), &model.Group{
		Id:    uuid.New(),
		Name:  req.Name,
		Order: len(groups),
	})
	if err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func (h *Handlers) PutGroup(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	req := new(PostGroupRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := h.Repo.PutGroup(c.Request().Context(), &model.Group{
		Id:   id,
		Name: req.Name,
	})
	if err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func (h *Handlers) DeleteGroup(c echo.Context) error {
	return c.JSON(200, "DeleteGroup")
}

func (h *Handlers) PutGroupUp(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	groups, err := h.Repo.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}

	for i, group := range groups {
		if group.Id == id {
			if i == 0 {
				return c.JSON(200, nil)
			}
			err = h.Repo.PutGroupOrder(c.Request().Context(), &model.Group{
				Id:    id,
				Order: group.Order - 1,
			})
			if err != nil {
				return err
			}
			err = h.Repo.PutGroupOrder(c.Request().Context(), &model.Group{
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

func (h *Handlers) PutGroupDown(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	groups, err := h.Repo.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}

	for i, group := range groups {
		if group.Id == id {
			if i == len(groups)-1 {
				return c.JSON(200, nil)
			}
			err = h.Repo.PutGroupOrder(c.Request().Context(), &model.Group{
				Id:    id,
				Order: group.Order + 1,
			})
			if err != nil {
				return err
			}
			err = h.Repo.PutGroupOrder(c.Request().Context(), &model.Group{
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

func (h *Handlers) PutGroupOrder(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))

	req := new(PutGroupOrderRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	newOrder := req.Order

	groups, err := h.Repo.GetGroups(c.Request().Context())
	if err != nil {
		return err
	}

	oldOrder := 0
	for _, group := range groups {
		if group.Id == id {
			oldOrder = group.Order
			break
		}
	}

	if oldOrder == newOrder {
		return c.JSON(200, nil)
	}

	if oldOrder < newOrder {
		for _, group := range groups {
			if group.Order <= newOrder && group.Order > oldOrder {
				err = h.Repo.PutGroupOrder(c.Request().Context(), &model.Group{
					Id:    group.Id,
					Order: group.Order - 1,
				})
				if err != nil {
					return err
				}
			}
		}
	} else {
		for _, group := range groups {
			if group.Order >= newOrder && group.Order < oldOrder {
				err = h.Repo.PutGroupOrder(c.Request().Context(), &model.Group{
					Id:    group.Id,
					Order: group.Order + 1,
				})
				if err != nil {
					return err
				}
			}
		}
	}

	err = h.Repo.PutGroupOrder(c.Request().Context(), &model.Group{
		Id:    id,
		Order: newOrder,
	})
	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}
