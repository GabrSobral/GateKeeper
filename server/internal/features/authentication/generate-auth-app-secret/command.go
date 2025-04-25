package generateauthappsecret

import "github.com/google/uuid"

type Command struct {
	ApplicationID uuid.UUID `json:"applicationId" validate:"required"`
	UserID        uuid.UUID `json:"userId" validate:"required"`
}
