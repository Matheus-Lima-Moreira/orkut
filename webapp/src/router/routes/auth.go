package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var authRoutes = []Route{
	{
		URI:       "/",
		Method:    http.MethodGet,
		Function:  controllers.LoadLoginPage,
		NeedsAuth: false,
	},
	{
		URI:       "/login",
		Method:    http.MethodGet,
		Function:  controllers.LoadLoginPage,
		NeedsAuth: false,
	},
	{
		URI:       "/login",
		Method:    http.MethodPost,
		Function:  controllers.Login,
		NeedsAuth: false,
	},
	{
		URI:       "/signup",
		Method:    http.MethodGet,
		Function:  controllers.LoadSignupPage,
		NeedsAuth: false,
	},
	{
		URI:       "/signup",
		Method:    http.MethodPost,
		Function:  controllers.Signup,
		NeedsAuth: false,
	},
	{
		URI:       "/logout",
		Method:    http.MethodGet,
		Function:  controllers.Logout,
		NeedsAuth: true,
	},
}
