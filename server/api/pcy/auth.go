package pcy

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/securecookie"
)

//ValidTokenPresent checks if the request has a valid cookie with the auth token
//If Must flag is set to true, the validator will fail if the token is not valid.
//Otherwise, the validator will fail if the token is present, so the request has
//already been validated
type ValidTokenPresent struct {
	CookieName string
	SecCookie  *securecookie.SecureCookie
	Must       bool
}

//ErrInvalidCookie error
var ErrInvalidCookie = errors.New("invalid cookie")

//ErrTokenPresent error means that the request is alreadi authenticated
var ErrTokenPresent = errors.New("token present")

//Validate satisfy HTTPPcy interface
func (pcy ValidTokenPresent) Validate(r *http.Request) error {

	err := func(r *http.Request) error {

		var cookie *http.Cookie
		var e error

		if cookie, e = r.Cookie(pcy.CookieName); e == nil || cookie != nil {
			value := make(map[string]string)
			if e = pcy.SecCookie.Decode(pcy.CookieName, cookie.Value, &value); e == nil {

				//TODO Validate cookie value here
				return nil

			}
			return ErrInvalidCookie
		}
		return http.ErrNoCookie
	}(r)

	if pcy.Must {
		return err
	}

	if err == http.ErrNoCookie {
		return nil
	}

	return ErrTokenPresent
}

//OnError handler
func (pcy ValidTokenPresent) OnError(w http.ResponseWriter, r *http.Request) {

	var cause string
	if pcy.Must {
		cause = "you must be logged in"
	} else {
		cause = "you are already logged in"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(map[string]string{"error": cause})
}
