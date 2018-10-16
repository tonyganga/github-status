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

// Endpoints to define the Route struct above.

var endpoints = Routes{
	Route{
		"GET",
		"/health",
		HealthCheckHandler,
	},
	Route{
		"GET",
		"/status",
		GithubStatusHandler,
	},
	Route{
		"GET",
		"/",
		IndexHandler,
	},
}
