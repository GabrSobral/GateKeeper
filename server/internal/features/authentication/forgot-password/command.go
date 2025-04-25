package forgotpassword

import "github.com/google/uuid"

type Command struct {
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
	Email         string    `json:"email" validate:"required,email"`
}
