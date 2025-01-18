package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homeRoutes = []Route{
	{
		URI:       "/home",
		Method:    http.MethodGet,
		Function:  controllers.LoadHomePage,
		NeedsAuth: true,
	},
}
