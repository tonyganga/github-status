package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// GithubStatusHandler is Github up?
func GithubStatusHandler(w http.ResponseWriter, r *http.Request) {

	req, err := http.Get("https://status.github.com/api/status.json")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(w, string(body))
}

// This is a fake healthcheck.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

// This handles requests to the index.
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello!")
}
