package entities

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationRole struct {
	ID            uuid.UUID
	ApplicationID uuid.UUID
	Name          string
	Description   *string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func NewApplicationRole(applicationID uuid.UUID, name string, description *string) *ApplicationRole {
	newID, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &ApplicationRole{
		ID:            newID,
		ApplicationID: applicationID,
		Name:          name,
		Description:   description,
		CreatedAt:     time.Now(),
		UpdatedAt:     nil,
	}
}
