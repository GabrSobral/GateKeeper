package removeapplication

import "github.com/google/uuid"

type Command struct {
	ApplicationID  uuid.UUID `json:"applicationId" validate:"required"`
	OrganizationID uuid.UUID `json:"organizationId" validate:"required"`
}
