package router

import "github.com/labstack/echo/v4"

func GetTags(c echo.Context) error {
	return c.JSON(200, "GetTags")
}

func PostTag(c echo.Context) error {
	return c.String(200, "PostTag")
}

func PutTag(c echo.Context) error {
	return c.String(200, "PutTag")
}

func DeleteTag(c echo.Context) error {
	return c.String(200, "DeleteTag")
}
