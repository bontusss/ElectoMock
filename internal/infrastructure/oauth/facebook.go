package oauth

import (
	"electomock/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type FacebookOauth struct {
	config *oauth2.Config
}

func NewfacebookOAuth(cfg config.OAuthConfig) *FacebookOauth {
	fbConfig := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes: []string{
			"email",
			"public_profile",
		},
		Endpoint: facebook.Endpoint,
	}
	return &FacebookOauth{config: fbConfig}
}

func (f FacebookOauth) GetAuthURL() string {
	return f.config.AuthCodeURL("state")
}
