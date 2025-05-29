package externalloginprovider

import "github.com/google/uuid"

type Command struct {
	ApplicationID uuid.UUID `json:"application" validate:"required"`
	Provider      string    `json:"provider" validate:"required"`
	ProviderKey   string    `json:"provider_key" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name" validate:"required"`
}
