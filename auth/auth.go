package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/coreos-inc/roller/context"
	"github.com/coreos-inc/roller/models"
)

type Auth struct {
	req *http.Request
	w   http.ResponseWriter
}

func getUser(req *http.Request, userName string) (user *models.AdminUser, ok bool, err error) {
	user = models.NewAdminUser(userName)
	store := context.Store(req)
	ok, err = store.Get(user)
	return
}

func SetUserLoggedIn(w http.ResponseWriter, req *http.Request,
	user *models.AdminUser) error {

	session, err := context.Session(req)
	if err != nil { // Usually an error decoding the session.
		log.Println(err) // There are many cases were we can't read the session, but should still try to move on
	}

	session.Values["user"] = user.User

	err = session.Save(req, w)
	if err != nil {
		log.Println(err)
	}

	return nil
}

// Returns true if auth succeeded.
func HawkAuth(w http.ResponseWriter, req *http.Request) bool {
	auther := &Auth{req, w}
	err := auther.DoHawkAuth()
	if err != nil {
		errorReply := map[string]interface{}{
			"error": map[string]interface{}{
				"code":    401,
				"message": err.Error(),
			},
		}
		w.WriteHeader(http.StatusUnauthorized)

		innerErr := json.NewEncoder(w).Encode(errorReply)
		if innerErr != nil {
			log.Println(err)
			log.Println(innerErr)
		}

		return false
	}

	return true
}

// Returns true if auth succeeded.
func SessionAuth(w http.ResponseWriter, req *http.Request) bool {
	session, err := context.Session(req)
	if err != nil { // Usually an error decoding the session.
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	userName, ok := session.Values["user"].(string)
	if session.IsNew || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	_, ok, err = getUser(req, userName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	if !ok {
		log.Println("Unauthorized user: ", userName)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func AuthWrapper(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		success := false
		if req.Header.Get("Authorization") != "" { // We are doing hawk auth.
			success = HawkAuth(w, req)
		} else {
			// Check for login session
			success = SessionAuth(w, req)
		}

		if success {
			// We didn't return an error, so serve req.
			handler.ServeHTTP(w, req)
		}
	}
}
