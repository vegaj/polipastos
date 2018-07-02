package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gobuffalo/pop"

	"github.com/polipastos/server/test"
)

var (
	//ErrInvalidRequest returned when a method was called with invalid parameters
	ErrInvalidRequest = errors.New("invalid request")

	db  *pop.Connection
	env test.Env
)

//Init mounts the database Depends on core
func Init(popconn *pop.Connection) {
	if popconn == nil {
		panic("Please, provide a valid database connection <nil>")
	}

	if db == nil {
		db = popconn
	}

}

//StartServer creates a new http server to attend the api endpoints
func StartServer() {

	r := mux.NewRouter()
	r.HandleFunc("/api/register", RegisterUserPost).Methods(http.MethodPost)

	log.Println("api server listening port 8080")
	log.Println(http.ListenAndServe(":8080", r))

}
