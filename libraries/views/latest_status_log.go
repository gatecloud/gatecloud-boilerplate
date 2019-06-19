package views

import (
	uuid "github.com/satori/go.uuid"
)

type LatestStatusLog struct {
	Model
	OwnerID     uuid.UUID
	OwnerType   string
	CompositeID uuid.UUID
	Status      string
	// Reference is plain text
	Reference string
	RoleName  string
	Email     string
	Name      string
	// Remark is cipher text written by admin
	Remark     string
	OperatorID uuid.UUID
}

// SQL script:

// DROP VIEW latest_status_logs;
// CREATE OR REPLACE VIEW latest_status_logs AS
// SELECT DISTINCT ON(owner_id) owner_id, id, created_at, updated_at, deleted_at, owner_type,
// composite_id, status, reference, role_name, email, name, remark, operator_id
// FROM status_logs
// ORDER BY owner_id asc, created_at desc
