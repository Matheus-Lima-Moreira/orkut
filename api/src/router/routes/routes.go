package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI       string
	Method    string
	Function  func(http.ResponseWriter, *http.Request)
	NeedsAuth bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoutes...)
	routes = append(routes, postRoutes...)

	for _, route := range routes {
		if route.NeedsAuth {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(
					middlewares.Authenticate(
						route.Function,
					),
				),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return router
}
