package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) GoogleLogin(c echo.Context) error {
	url := h.issueLoginURL()
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *Handlers) GoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != googleState {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get token")
	}

	code := c.QueryParam("code")
	token, err := h.getTokenFromCode(code)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get token")
	}

	email, err := h.getGmailFromToken(token)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get gmail address")
	}
	if !h.isWhiteList(email) {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get gmail address")
	}

	// TODO: save refresh token (token.RefreshToken) to db

	err = h.login(c, token.AccessToken)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to login")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *Handlers) login(c echo.Context, googleToken string) error {
	token, err := h.Repo.IssueToken(c.Request().Context(), googleToken)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    token.TokenID,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HttpOnly: true,
	})
	return nil
}

func (h *Handlers) isWhiteList(email string) bool {
	for _, v := range h.WhiteList {
		if v == email {
			return true
		}
	}
	return false
}
