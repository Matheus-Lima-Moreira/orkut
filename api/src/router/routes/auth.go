package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoutes = []Route{
	{
		URI:    "/auth/login",
		Method: http.MethodPost,
		Function: controllers.Login,
		NeedsAuth: false,
	},
}
