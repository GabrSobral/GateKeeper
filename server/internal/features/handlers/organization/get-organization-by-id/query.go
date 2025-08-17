package getorganizationbyid

import "github.com/google/uuid"

type Query struct {
	OrganizationID uuid.UUID
}
