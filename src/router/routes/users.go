package routes

import (
	"golang-api/src/controllers"
	"net/http"
)

var usersRoute = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Func:                   controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Func:                   controllers.GetUsers,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Func:                   controllers.GetUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Func:                   controllers.UpdateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Func:                   controllers.DeleteUser,
		AuthenticationRequired: false,
	},
}
