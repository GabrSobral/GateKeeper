package changepassword

import "github.com/google/uuid"

type Command struct {
	ChangePasswordCode string    `json:"changePasswordCode" validate:"required,min=64,max=64"`
	ApplicationID      uuid.UUID `json:"applicationID" validate:"required"`
	UserID             uuid.UUID `json:"userID" validate:"required"`
	NewPassword        string    `json:"newPassword" validate:"required,min=8,max=64"`
}
