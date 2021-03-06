package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TestTask",
		"POST",
		"/api/v1/testtask",
		TestTask,
	},
	Route{
		"RunTool",
		"POST",
		"/api/v1/runtool",
		RunTool,
	},
}
