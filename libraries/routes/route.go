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

func (r Route) Register(action string, sr *SharedResource) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		languageCode := ctx.Request.Header.Get("X-Language-Code")
		ptr := reflect.New(r.Controller)
		methodInit := ptr.MethodByName("Init")
		if methodInit.IsValid() {
			args := make([]reflect.Value, 6)
			args[0] = reflect.ValueOf(sr.DB)
			args[1] = reflect.ValueOf(r.Name)
			args[2] = reflect.ValueOf(languageCode)
			args[3] = reflect.ValueOf(sr.Validate)
			args[4] = reflect.ValueOf(r.Model)
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
