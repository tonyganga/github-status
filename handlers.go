package main

import (
	"fmt"
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
