package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postRoutes = []Route{
	{
		URI:       "/posts",
		Method:    http.MethodPost,
		Function:  controllers.CreatePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}/edit",
		Method:    http.MethodGet,
		Function:  controllers.LoadEditPostPage,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}",
		Method:    http.MethodPut,
		Function:  controllers.EditPost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}/like",
		Method:    http.MethodPost,
		Function:  controllers.LikePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}/dislike",
		Method:    http.MethodPost,
		Function:  controllers.DislikePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}",
		Method:    http.MethodDelete,
		Function:  controllers.DeletePost,
		NeedsAuth: true,
	},
}
