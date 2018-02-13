package main

import "net/http"

// Route is the struct we use to define routes in the mux router.
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is how we add routers.
type Routes []Route
