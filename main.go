package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// init router
	r := mux.NewRouter().StrictSlash(true)

	// create endpoints defined in routes.go
	for _, item := range endpoints {
		r.
			Methods(item.Method).
			Path(item.Pattern).
			HandlerFunc(item.HandlerFunc)
	}

	// init server
	s := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
