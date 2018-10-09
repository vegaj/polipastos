package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gobuffalo/uuid"
	"github.com/gorilla/sessions"
	"github.com/polipastos/server/core"
)

//LoginUser endpoint stores a uuid encrypted cookie in the client
func LoginUser(w http.ResponseWriter, r *http.Request) {

	var nulableid uuid.NullUUID
	var session *sessions.Session
	var status int

	status = http.StatusBadRequest
	m := make(url.Values)
	err := func(r *http.Request) error {

		var e error
		session, e = store.Get(r, sesName)
		if e != nil && !session.IsNew {
			return e
		}

		pair, e := retrieveCredentials(r)
		if e != nil {
			return e
		}

		id, e := core.PairAuthentication(pair[0], pair[1])
		if e != nil {
			status = http.StatusForbidden
			return e
		}

		nulableid.UUID = id
		nulableid.Valid = true

		return nil
	}(r)

	if err == nil {
		session.Values[sesUUIDKey] = nulableid.UUID.Bytes()
		m.Add("ok", "true")
		m.Add("uuid", nulableid.UUID.String())
		status = http.StatusOK

		token, _ := uuid.NewV4()
		value := map[string]string{
			authCookieName: token.String(),
		}

		if encoded, err := authSecCookie.Encode(authCookieName, value); err == nil {
			http.SetCookie(w, NewAuthCookie(authCookieName, encoded))
		} else {
			panic(err)
		}

		if err := session.Save(r, w); err != nil {
			panic(err)
		}

	} else {
		m.Add("error", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(m)
}

func retrieveCredentials(r *http.Request) ([]string, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	uname := r.FormValue("username")
	passw := r.FormValue("password")

	if uname == "" || passw == "" {
		return nil, ErrInvalidRequest
	}

	return []string{uname, passw}, nil
}

//NewAuthCookie configures a cookie with the given value with the standard configuration
func NewAuthCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(authHeartBeatDur),
		HttpOnly: true,
		Secure:   false, //TODO set to true when changin form HTTP to HTTPS
	}
}

//KeepAliveCookie increments the Expiration time
func KeepAliveCookie(c *http.Cookie) {
	c.Expires = time.Now().Add(authHeartBeatDur)
}
