package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI       string
	Method    string
	Function  func(http.ResponseWriter, *http.Request)
	NeedsAuth bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, signupRoutes...)
	routes = append(routes, homeRoutes...)
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

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", fileServer))

	return router
}
