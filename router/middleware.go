package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getTokenID(c echo.Context) string {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		return "-"
	}
	return cookie.Value
}

func (h *Handlers) EnsureAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		permission, err := h.Repo.CertificateToken(c.Request().Context(), getTokenID(c))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to get permission: %v", err))
		}

		if !permission {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		return next(c)
	}
}
