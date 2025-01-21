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
	{
		URI:       "/profile",
		Method:    http.MethodGet,
		Function:  controllers.LoadUserLoggedInProfilePage,
		NeedsAuth: true,
	},
	{
		URI:       "/edit-profile",
		Method:    http.MethodGet,
		Function:  controllers.LoadEditProfilePage,
		NeedsAuth: true,
	},
	{
		URI:       "/edit-profile",
		Method:    http.MethodPut,
		Function:  controllers.EditProfile,
		NeedsAuth: true,
	},
	{
		URI:       "/update-password",
		Method:    http.MethodGet,
		Function:  controllers.LoadUpdatePasswordPage,
		NeedsAuth: true,
	},
	{
		URI:       "/update-password",
		Method:    http.MethodPut,
		Function:  controllers.UpdatePassword,
		NeedsAuth: true,
	},
	{
		URI:       "/delete-account",
    Method:    http.MethodDelete,
    Function:  controllers.DeleteAccount,
    NeedsAuth: true,
	},
}
