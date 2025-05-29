package resetpassword

import "github.com/google/uuid"

type Command struct {
	PasswordResetToken string    `json:"passwordResetToken" validate:"required"`
	PasswordResetId    uuid.UUID `json:"passwordResetId" validate:"required"`
	NewPassword        string    `json:"newPassword" validate:"required"`
	ApplicationID      uuid.UUID `json:"applicationId" validate:"required"`
}
