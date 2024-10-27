package main

import (
	"log"
	"net/http"
)

type application struct {
	config config
}

type config struct {
	addr string

}

func (app *application) run() error{

	mux := http.NewServeMux()

	srv := http.Server{
		Addr: app.config.addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", app.config.addr)
	return srv.ListenAndServe()
}