package routes

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// Route stores the API route information
type Route struct {
	Name       string
	Controller reflect.Type
	Model      interface{}
}

// Register registers the API to route
func (r *Route) Register(action string, sr *SharedResource) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ptr := reflect.New(r.Controller)
		methodInit := ptr.MethodByName("Init")
		if methodInit.IsValid() {
			args := make([]reflect.Value, 4)
			args[0] = reflect.ValueOf(sr.DB)
			args[1] = reflect.ValueOf(sr.Validator)
			args[2] = reflect.ValueOf(r.Model)
			args[3] = reflect.ValueOf(sr.RedisClient)
			methodInit.Call(args)
		}

		m := ptr.MethodByName(action)
		if m.IsValid() {
			args := make([]reflect.Value, 1)
			args[0] = reflect.ValueOf(ctx)
			m.Call(args)
		}
	}
}
