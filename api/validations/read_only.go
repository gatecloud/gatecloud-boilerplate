package validations

import (
	"gatecloud-boilerplate/api/models"
	"reflect"

	validator "gopkg.in/go-playground/validator.v8"
)

func ReadOnlyValidation(v *validator.Validate, sl *validator.StructLevel) {
	switch v := sl.CurrentStruct.Interface().(type) {
	case models.TestAPI:
		if v.Number != "" {
			sl.ReportError(reflect.ValueOf(v), "Number", "read_only", "read_only")
		}
	}
}
