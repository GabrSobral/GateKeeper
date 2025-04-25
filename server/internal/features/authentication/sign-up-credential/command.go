package signupcredential

import "github.com/google/uuid"

type Command struct {
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
	DisplayName   string    `json:"displayName"`
	FirstName     string    `json:"firstName" validate:"required"`
	LastName      string    `json:"lastName" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	Password      string    `json:"password" validate:"required"`
}
