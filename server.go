package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func server() {
	// init router
	r := mux.NewRouter().StrictSlash(true)
	for _, item := range routes {
		r.
			Methods(item.Method).
			Path(item.Pattern).
			HandlerFunc(item.HandlerFunc)
	}

	// init server
	s := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())

	// handle index
	http.Handle("/", r)
}
