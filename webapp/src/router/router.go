package router

import (
	"github.com/gorilla/mux"
	"webapp/src/router/routes"
)

func Generate() *mux.Router {

	r := mux.NewRouter()

	return routes.Configure(r)
}
