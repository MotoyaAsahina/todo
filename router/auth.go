package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/MotoyaAsahina/todo/model"
)

func GoogleLogin(c echo.Context) error {
	url := issueLoginURL()
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != googleState {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get token")
	}

	code := c.QueryParam("code")
	token, err := getTokenFromCode(code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get token")
	}

	email, err := getGmailFromToken(token)
	if !isWhiteList(email) {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get gmail address")
	}

	// TODO: save refresh token (token.RefreshToken) to db

	err = login(c, token.AccessToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to login")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func login(c echo.Context, googleToken string) error {
	token, err := model.IssueToken(c.Request().Context(), googleToken)
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

func isWhiteList(email string) bool {
	for _, v := range whiteList {
		if v == email {
			return true
		}
	}
	return false
}
