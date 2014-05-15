package auth

import (
	"github.com/coreos-inc/roller/third_party/code.google.com/p/goauth2/oauth"
	"github.com/coreos-inc/roller/third_party/github.com/gorilla/sessions"
)

type Config struct {
	OauthConfig  *oauth.Config
	SessionStore *sessions.CookieStore
}

func NewOauthConfig(clientId, clientSecret, redirectURL string) *oauth.Config {
	return &oauth.Config{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scope:        "openid email",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://accounts.google.com/o/oauth2/token",
	}
}

func NewConfig(sessionSecret string) *Config {
	return &Config{
		OauthConfig: &oauth.Config{
			Scope:    "openid email",
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
		SessionStore: sessions.NewCookieStore([]byte(sessionSecret)),
	}
}
