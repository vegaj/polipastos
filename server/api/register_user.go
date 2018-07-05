package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gobuffalo/uuid"

	"github.com/polipastos/server/core"
)

//RegisterUserPost endpoint
func RegisterUserPost(w http.ResponseWriter, r *http.Request) {

	resp := json.NewEncoder(w)
	var ok = true
	m := make(url.Values)

	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		log.Println("RegisterUserPost - ", err)
		m.Set("error", ErrInvalidRequest.Error())
		ok = false
	}

	var username, password string
	username = r.FormValue("username")
	password = r.FormValue("password")

	if ok && (username == "" || password == "") {
		w.WriteHeader(http.StatusBadRequest)
		m.Set("error", ErrInvalidRequest.Error())
		ok = false
	}

	var id uuid.UUID

	if ok {
		id, err = core.RegisterUser(username, password)
		if err != nil {
			m.Set("error", err.Error())
			ok = false
		}
	}

	if ok {

		//Set cookies. If this fails, the user will have to perform a manual login
		token, _ := uuid.NewV4()
		if encoded, err := authSecCookie.Encode(authCookieName, map[string]string{authCookieName: token.String()}); err == nil {
			http.SetCookie(w, NewAuthCookie(authCookieName, encoded))
		} else {
			log.Println("Could not set cookie", err)
		}

		//Set session uuid cookie. If this fails, the uuid must be retrieved
		if sess, err := store.Get(r, sesName); err == nil {
			sess.Values[sesUUIDKey] = id.Bytes()
			sess.Save(r, w)
		} else {
			log.Println("Could not store session cookies", err)
		}

		m.Set("ok", "true")
		m.Set("uuid", id.String())
		w.WriteHeader(http.StatusOK)
	}
	resp.Encode(m)
}
