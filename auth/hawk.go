package auth

import (
	"crypto/sha256"

	"github.com/coreos-inc/roller/third_party/github.com/tent/hawk-go"
)

var DefaultHawkHasher = sha256.New

func (a *Auth) HawkLookupCreds(c *hawk.Credentials) error {
	u, ok, err := getUser(a.req, c.ID)
	if !ok || err != nil {
		return &hawk.CredentialError{hawk.UnknownID, c}
	}

	c.Key = u.Token
	c.Hash = DefaultHawkHasher
	return nil
}

func (a *Auth) DoHawkAuth() error {
	auth, err := hawk.NewAuthFromRequest(a.req, a.HawkLookupCreds, nil)

	if err != nil {
		return err
	}

	return auth.Valid()
}
