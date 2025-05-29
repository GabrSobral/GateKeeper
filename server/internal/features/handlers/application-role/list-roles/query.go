package listroles

import "github.com/google/uuid"

type Query struct {
	ApplicationID  uuid.UUID `json:"applicationId" validate:"required,uuid"`
	OrganizationID uuid.UUID `json:"organizationId" validate:"required,uuid"`
}
