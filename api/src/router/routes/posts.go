package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:       "/posts",
		Method:    http.MethodPost,
		Function:  controllers.CreatePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts",
		Method:    http.MethodGet,
		Function:  controllers.ListPosts,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}",
		Method:    http.MethodGet,
		Function:  controllers.ShowPost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}",
		Method:    http.MethodPut,
		Function:  controllers.UpdatePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}",
		Method:    http.MethodDelete,
		Function:  controllers.DeletePost,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}/posts",
		Method:    http.MethodGet,
		Function:  controllers.GetPostsByUser,
		NeedsAuth: true,
	},
}
