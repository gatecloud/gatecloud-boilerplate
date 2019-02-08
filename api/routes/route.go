package routes

import (
	"gatecloud-boilerplate/api/controllers"
	"gatecloud-boilerplate/api/models"
	"reflect"

	libRoute "github.com/gatecloud/webservice-library/route"
)

// InitRoute inits the route group
func InitRoute() []libRoute.Route {
	r := []libRoute.Route{
		libRoute.Route{
			Name:       "TestAPI",
			Controller: reflect.TypeOf(controllers.TestAPIControl{}),
			Model:      &models.TestAPI{},
		},
	}
	return r
}
