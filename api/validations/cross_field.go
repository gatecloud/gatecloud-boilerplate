package validations

import (
	"gatecloud-boilerplate/api/models"
	"reflect"
	"time"

	"gopkg.in/go-playground/validator.v8"
)

func CrossFieldValidation(v *validator.Validate, sl *validator.StructLevel) {
	switch v := sl.CurrentStruct.Interface().(type) {
	case models.TestAPI:
		if v.From == time.Duration(0) && v.To == time.Duration(0) {
			sl.ReportError(reflect.ValueOf(v.From), "From", "at_least_one_required", "at_least_one_required")
			sl.ReportError(reflect.ValueOf(v.To), "To", "at_least_one_required", "at_least_one_required")
		}
	}
}
