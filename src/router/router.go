package router

import (
	"github.com/gorilla/mux"
	"golang-api/src/router/routes"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
