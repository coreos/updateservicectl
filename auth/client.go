package auth

import (
	"net/http"

	"github.com/coreos-inc/roller/third_party/github.com/tent/hawk-go"
)

type HawkRoundTripper struct {
	User  string
	Token string
}

func (t *HawkRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	creds := &hawk.Credentials{
		ID:   t.User,
		Key:  t.Token,
		Hash: DefaultHawkHasher,
	}

	auth := hawk.NewRequestAuth(req, creds, 0)

	req.Header.Set("Authorization", auth.RequestHeader())
	return http.DefaultTransport.RoundTrip(req)
}
