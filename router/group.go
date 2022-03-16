package router

import "github.com/labstack/echo/v4"

func GetGroups(c echo.Context) error {
	return c.JSON(200, "GetGroups")
}

func PostGroup(c echo.Context) error {
	return c.JSON(200, "PostGroup")
}

func PutGroup(c echo.Context) error {
	return c.JSON(200, "PutGroup")
}

func DeleteGroup(c echo.Context) error {
	return c.JSON(200, "DeleteGroup")
}
