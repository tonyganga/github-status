package main

var routes = Routes{
	Route{
		"GET",
		"/health",
		HealthCheckHandler,
	},
}
