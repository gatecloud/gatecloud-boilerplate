package models

import (
	"timingniao_wlx_api/libraries/views"

	uuid "github.com/satori/go.uuid"
)

type StatusLog struct {
	Model
	OwnerID     uuid.UUID `validate:"required" gorm:"type:uuid;not null"`
	OwnerType   string    `json:"-"`
	CompositeID uuid.UUID
	Status      string `validate:"required" type:"not null;"`
	// Reference is plain text
	Reference string
	RoleName  string
	Email     string
	Name      string
	// Remark is cipher text written by admin
	Remark     string
	OperatorID uuid.UUID `gorm:"type:uuid"`
}

func (s *StatusLog) ConvertToLatest() views.LatestStatusLog {
	return views.LatestStatusLog{
		Model: views.Model{
			ID:        s.Model.ID,
			CreatedAt: s.Model.CreatedAt,
			UpdatedAt: s.Model.UpdatedAt,
			DeletedAt: s.DeletedAt,
		},
		OwnerID:     s.OwnerID,
		OwnerType:   s.OwnerType,
		CompositeID: s.CompositeID,
		Status:      s.Status,
		Reference:   s.Reference,
		RoleName:    s.RoleName,
		Email:       s.Email,
		Name:        s.Name,
		Remark:      s.Remark,
		OperatorID:  s.OperatorID,
	}
}
