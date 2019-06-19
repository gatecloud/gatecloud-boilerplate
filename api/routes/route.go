package routes

import (
	"gatecloud-boilerplate/api/controllers"
	"gatecloud-boilerplate/api/models"
	"reflect"

	libRoute "github.com/gatecloud/webservice-library/route"
)

var (
	RouteMap map[string][]libRoute.Route
)

// InitRoute inits the route group
func init() {
	RouteMap = make(map[string][]libRoute.Route)
	RouteMap["api"] = []libRoute.Route{
		libRoute.Route{
			Name:       "TestAPI",
			Controller: reflect.TypeOf(controllers.TestAPIControl{}),
			Model:      &models.TestAPI{},
		},
	}

	RouteMap["test"] = []libRoute.Route{
		libRoute.Route{
			Name:       "TestAPI",
			Controller: reflect.TypeOf(controllers.TestAPIControl{}),
			Model:      &models.TestAPI{},
		},
	}

}
