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

	//Should be checked by the router
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		m.Set("error", "invalid request method")
		resp.Encode(m)
		return
	}

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
		m.Set("ok", "true")
		m.Set("uuid", id.String())
		w.WriteHeader(http.StatusOK)
	}
	resp.Encode(m)
}
