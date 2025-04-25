package getapplicationauthdata

import "github.com/google/uuid"

type Query struct {
	ApplicationID uuid.UUID `json:"applicationId" validate:"required"`
}
