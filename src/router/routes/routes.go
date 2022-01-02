package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route Represente the structure
type Route struct {
	URI                    string
	Method                 string
	Func                   func(w http.ResponseWriter, r *http.Request)
	AuthenticationRequired bool
}

// Config Configure all routes at api
func Config(r *mux.Router) *mux.Router {
	routes := usersRoute

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
