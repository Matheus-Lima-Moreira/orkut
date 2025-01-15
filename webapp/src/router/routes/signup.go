package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var signupRoutes = []Route{
	{
		URI:       "/signup",
		Method:    http.MethodGet,
		Function:  controllers.LoadSignupPage,
		NeedsAuth: false,
	},
	{
		URI:       "/signup",
    Method:    http.MethodPost,
    Function:  controllers.CreateUser,
    NeedsAuth: false,
	},
}
