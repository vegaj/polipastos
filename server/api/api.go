package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/polipastos/server/api/pcy"

	"github.com/gorilla/securecookie"

	"github.com/gorilla/mux"

	"github.com/gobuffalo/pop"
	"github.com/gorilla/sessions"

	"github.com/polipastos/server/test"
)

var (
	//ErrInvalidRequest returned when a method was called with invalid parameters
	ErrInvalidRequest = errors.New("invalid request")
	//ErrAlreadyLoggedIn a login request was made when the caller was already logged in
	ErrAlreadyLoggedIn = errors.New("already logged in")
	//ErrCookieExpired .
	ErrCookieExpired = errors.New("the cookie has expired")

	db  *pop.Connection
	env test.Env

	store       = sessions.NewCookieStore(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(24))
	sesUUIDKey  = "polipastos-uuid-key"
	sesName     = "polipastos-ses"
	sesDuration time.Duration

	authCookieName   = "polipastos-auth"
	authSecCookie    = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(24))
	authHeartBeatDur time.Duration
)

//Init mounts the database Depends on core
func Init(popconn *pop.Connection) {
	if popconn == nil {
		panic("Please, provide a valid database connection <nil>")
	}

	if db == nil {
		db = popconn
	}

	authHeartBeatDur, _ = time.ParseDuration("1m")
	sesDuration, _ = time.ParseDuration("1m")

	store.Options.MaxAge = int(sesDuration.Seconds())
	store.Options.Path = "/"
}

//StartServer creates a new http server to attend the api endpoints
func StartServer() {

	var validToken pcy.ValidTokenPresent
	validToken.CookieName = authCookieName
	validToken.SecCookie = authSecCookie
	validToken.Must = true

	secretHandler := pcy.NewHandler(dummySecret, validToken.OnError, validToken)

	validToken.Must = false
	loginHandler := pcy.NewHandler(LoginUser, validToken.OnError, validToken)
	registerHandler := pcy.NewHandler(RegisterUserPost, validToken.OnError, validToken)

	r := mux.NewRouter()
	r.HandleFunc("/api/register", registerHandler.ServeHTTP).Methods(http.MethodPost)
	r.HandleFunc("/api/login", loginHandler.ServeHTTP).Methods(http.MethodPost)
	r.HandleFunc("/api/secret", secretHandler.ServeHTTP).Methods(http.MethodGet)

	log.Println("api server listening port 8080")
	log.Println(http.ListenAndServe(":8080", r))

}

func dummySecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("this is my secret")
}
