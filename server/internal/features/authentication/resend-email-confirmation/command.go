package resendemailconfirmation

import "github.com/google/uuid"

type Command struct {
	ApplicationID uuid.UUID `json:"applicationId"`
	Email         string    `json:"email"`
}
