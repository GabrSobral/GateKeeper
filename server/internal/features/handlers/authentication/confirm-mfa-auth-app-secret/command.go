package confirmmfaauthappsecret

import "github.com/google/uuid"

type Command struct {
	UserID         uuid.UUID `json:"userId" validate:"required"`
	MfaAuthAppCode string    `json:"mfaAuthAppCode" validate:"required"`
}
