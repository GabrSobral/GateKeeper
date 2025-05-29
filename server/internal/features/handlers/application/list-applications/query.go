package listapplications

import "github.com/google/uuid"

type Query struct {
	OrganizationID uuid.UUID `json:"organizationId" validate:"required"`
}
