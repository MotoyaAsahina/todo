package middleware

import (
	"fmt"
	"github.com/MotoyaAsahina/todo/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getTokenID(c echo.Context) string {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		return "-"
	}
	return cookie.Value
}

func EnsureAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		permission, err := model.CertificateToken(c.Request().Context(), getTokenID(c))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to get permission: %v", err))
		}

		if !permission {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		return next(c)
	}
}
