package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		isAuthRequired: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.GetAllUsers,
		isAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Function:       controllers.GetUserById,
		isAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		isAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		isAuthRequired: false,
	},
}
