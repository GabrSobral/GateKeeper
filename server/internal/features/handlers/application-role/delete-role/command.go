package deleterole

import "github.com/google/uuid"

type Command struct {
	RoleID        uuid.UUID `json:"roleId" validate:"required,uuid"`
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
}
