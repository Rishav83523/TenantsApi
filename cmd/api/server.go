package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func (app *application) Serve() error {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", app.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Server started on port %d", app.port)
	return server.ListenAndServe()
}