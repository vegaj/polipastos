package main

import (
	"log"

	"github.com/gobuffalo/pop"
	"github.com/polipastos/server/api"
	"github.com/polipastos/server/core"
)

func main() {

	db, err := pop.Connect("development")
	if err != nil {
		panic(err)
	}

	api.Init(db)
	core.Init(db)

	/*
		r := mux.NewRouter()
		r.HandleFunc("/api/register", api.RegisterUserPost)

		http.ListenAndServe(":8080", r)
	*/

	api.StartServer()
	log.Println("----")
}
