package createsecret

import (
	"time"

	"github.com/google/uuid"
)

type Command struct {
	ApplicationID uuid.UUID  `json:"applicationId" validate:"required,uuid"`
	Name          string     `json:"name" validate:"required"`
	ExpiresAt     *time.Time `json:"expiresAt"`
}

type RequestBody struct {
	Name      string     `json:"name" validate:"required"`
	ExpiresAt *time.Time `json:"expiresAt"`
}
