package deletesecret

import (
	"github.com/google/uuid"
)

type Command struct {
	SecretID      uuid.UUID `json:"secretId" validate:"required,uuid"`
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
}
