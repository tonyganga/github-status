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

// The actual routes.

var myRoutes = Routes{
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
