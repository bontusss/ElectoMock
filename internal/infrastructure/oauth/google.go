package oauth

import (
	"context"
	"electomock/config"
	"electomock/internal/domain/models"
	"electomock/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
)

type GoogleOauth struct {
	config *oauth2.Config
}

func NewGoogleOAuth(cfg config.OAuthConfig) *GoogleOauth {
	gConfig := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return &GoogleOauth{config: gConfig}
}

func (g GoogleOauth) GetAuthURL() string {
	return g.config.AuthCodeURL("state")
}

func (g GoogleOauth) HandleGoogleCallback(code string) (*models.User, error) {
	token, err := g.config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("%s: failed to exchange code for token.\n", utils.GOOGLE_AUTH_ERROR)
		return nil, errors.New("google authentication failed")
	}

	client := g.config.Client(context.Background(), token)

	userInfoUrl := "https://www.googleapis.com/oauth2/v2/userinfo"

	resp, err := client.Get(userInfoUrl)
	if err != nil {
		fmt.Printf("%s: failed to retrieve user info from google\n", utils.GOOGLE_AUTH_ERROR)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("%s: %v", utils.GOOGLE_AUTH_ERROR, err)
		}
	}(resp.Body)

	var user *models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Printf("%s: failed to parse user info", utils.GOOGLE_AUTH_ERROR)
		return nil, errors.New("google authentication failed")
	}

	return user, nil
}
