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
    NeedsAuth: false,
	},
	{
		URI:       "/users/{userId}",
    Method:    http.MethodGet,
    Function:  controllers.ShowUser,
    NeedsAuth: false,
	},
	{
		URI:       "/users/{userId}",
    Method:    http.MethodPut,
    Function:  controllers.UpdateUser,
    NeedsAuth: false,
	},
	{
		URI:       "/users/{userId}",
    Method:    http.MethodDelete,
    Function:  controllers.DeleteUser,
    NeedsAuth: false,
	},
}