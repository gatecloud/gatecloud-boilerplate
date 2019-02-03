package routes

import (
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

type SharedResource struct {
	DB       *gorm.DB
	Validate *validator.Validate
}
