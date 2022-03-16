package router

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const googleState = "todo_auth_state"

var (
	config    *oauth2.Config
	whiteList []string
)

func SetupGoogleOauth2() {
	config = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	whiteList = strings.Split(os.Getenv("WHITE_LIST"), ",")
}

func issueLoginURL() string {
	return config.AuthCodeURL(googleState, oauth2.AccessTypeOffline)
}

func getTokenFromCode(code string) (*oauth2.Token, error) {
	return config.Exchange(context.Background(), code)
}

func getGmailFromToken(token *oauth2.Token) (string, error) {
	client := config.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	user := struct {
		Email string `json:"email"`
	}{}
	err = dec.Decode(&user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
