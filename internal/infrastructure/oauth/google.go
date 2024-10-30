package oauth

import (
	"electomock/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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