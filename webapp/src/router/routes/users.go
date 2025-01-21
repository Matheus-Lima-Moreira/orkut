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
	{
		URI:       "/users/{userId}/unfollow",
		Method:    http.MethodPost,
		Function:  controllers.Unfollow,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}/follow",
		Method:    http.MethodPost,
		Function:  controllers.Follow,
		NeedsAuth: true,
	},
}
