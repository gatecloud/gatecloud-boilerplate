package routes

import (
	"gatecloud-boilerplate/api/controllers"
	"gatecloud-boilerplate/api/models"
	"gatecloud-boilerplate/libraries/routes"
	"reflect"
)

// InitRoute inits the route group
func InitRoute() []routes.Route {
	r := []routes.Route{
		routes.Route{
			Name:       "TestAPI",
			Controller: reflect.TypeOf(controllers.TestAPIControl{}),
			Model:      &models.TestAPI{},
		},
	}
	return r
}
