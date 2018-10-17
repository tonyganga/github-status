package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Github struct {
	Status      string `json:"status"`
	LastUpdated string `json:"created_on"`
	LastMessage string `json:"body"`
	Available   bool
}

func GithubStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Read body for github.com/api/last-message.json
	res, err := http.Get("https://status.github.com/api/last-message.json")
	log.Printf("GET /api/last-message.json")
	if err != nil {
		log.Fatal(err)
	}
	status, err := ioutil.ReadAll(res.Body)
	log.Printf("Closing body of request")
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Unmarshal the body of the request into the Github struct
	var github Github
	json.Unmarshal(status, &github)
	log.Printf("Unmarshaling last-message.json")
	if err != nil {
		log.Printf("Unable to unmarshal last-message.json")
		http.Error(w, err.Error(), 500)
	}
	// Determine if Github is good then set Available to true or false.
	if github.Status == "good" {
		log.Printf("Github status is good.")
		GithubIsAvailable(&github)
		output, err := json.Marshal(github)
		if err != nil {
			log.Printf("Unable to marshal last-message.json")
			http.Error(w, err.Error(), 500)
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	} else {
		log.Printf("Github returned an outage state.")
		GithubIsNotAvailable(&github)
		output, err := json.Marshal(github)
		if err != nil {
			log.Printf("Unable to marshal last-message.json")
			http.Error(w, err.Error(), 500)
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}

func GithubIsAvailable(g *Github) {
	g.Available = true
}

func GithubIsNotAvailable(g *Github) {
	g.Available = false
}

// This is a fake healthcheck.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

// This handles requests to the index.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
