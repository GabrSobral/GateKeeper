package verifymfa

import "github.com/google/uuid"

type Command struct {
	Code          string    `json:"code" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	ApplicationID uuid.UUID `json:"applicationId" validate:"required"`
}
