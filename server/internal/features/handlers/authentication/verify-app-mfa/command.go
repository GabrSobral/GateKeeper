package verifyappmfa

import "github.com/google/uuid"

type Command struct {
	Code          string     `json:"code" validate:"required"`
	Email         string     `json:"email" validate:"required,email"`
	MfaID         *uuid.UUID `json:"mfaId" validate:"required,uuid"`
	ApplicationID uuid.UUID  `json:"applicationId" validate:"required"`
}
