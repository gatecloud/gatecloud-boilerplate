package models

type Role struct {
	Model
	Name        string `validate:"required" gorm:"unique"`
	DisplayName string `validate:"required" gorm:"unique"`
	OwnerType   string
	Description string
}
