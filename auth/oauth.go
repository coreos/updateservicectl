package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/coreos-inc/roller/context"
	"github.com/coreos-inc/roller/models"

	"github.com/coreos-inc/roller/third_party/code.google.com/p/goauth2/oauth"
)

func OauthRedirect(w http.ResponseWriter, req *http.Request) {
	config := context.OAuthConfig(req)

	url := config.AuthCodeURL("")
	http.Redirect(w, req, url, 302)
}

func OAuthLogin(req *http.Request) (user *models.AdminUser, ok bool, err error) {
	code := req.URL.Query().Get("code")
	if code == "" {
		err = errors.New("OAuthLogin: code was not in query string.")
		return
	}

	config := context.OAuthConfig(req)

	transport := &oauth.Transport{Config: config}

	token, err := transport.Exchange(code)
	if err != nil {
		return
	}
	transport.Token = token

	r, err := transport.Client().Get("https://www.googleapis.com/userinfo/v2/me")
	if err != nil {
		return
	}
	defer r.Body.Close()

	var userData struct {
		Email string `json:"email"`
	}

	err = json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		return
	}

	user, ok, err = getUser(req, userData.Email)
	if err != nil {
		return
	}

	if !ok {
		return
	}

	return
}
