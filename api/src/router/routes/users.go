package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:       "/users",
		Method:    http.MethodPost,
		Function:  controllers.CreateUser,
		NeedsAuth: false,
	},
	{
		URI:       "/users",
		Method:    http.MethodGet,
		Function:  controllers.ListUsers,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodGet,
		Function:  controllers.ShowUser,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodPut,
		Function:  controllers.UpdateUser,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodDelete,
		Function:  controllers.DeleteUser,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}/follow",
		Method:    http.MethodPost,
		Function:  controllers.FollowUser,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}/unfollow",
		Method:    http.MethodPost,
		Function:  controllers.UnfollowUser,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}/followers",
		Method:    http.MethodGet,
		Function:  controllers.GetFollowers,
		NeedsAuth: true,
	},
}
