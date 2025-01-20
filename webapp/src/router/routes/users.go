package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:       "/search-users",
		Method:    http.MethodGet,
		Function:  controllers.LoadUsersPage,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodGet,
		Function:  controllers.LoadUserProfilePage,
		NeedsAuth: true,
	},
}
