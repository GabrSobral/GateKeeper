package getapplicationbyid

import "github.com/google/uuid"

type Query struct {
	ApplicationID  uuid.UUID `json:"applicationId" validate:"required"`
	OrganizationID uuid.UUID `json:"organizationId" validate:"required"`
}
