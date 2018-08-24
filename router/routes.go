package router

import (
	"net/http"
	"go-server/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAllUsers",
		"GET",
		"/users",
		controllers.GetAllUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/users/{userId}",
		controllers.GetUser,
	},
	Route{
		"PostUser",
		"POST",
		"/users",
		controllers.PostUser,
	},
}