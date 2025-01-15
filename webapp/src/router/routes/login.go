package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRoutes = []Route{
	{
		URI: "/",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		NeedsAuth: false,
	},
	{
		URI: "/login",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		NeedsAuth: false,
	},
}