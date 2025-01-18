package entities

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID            uuid.UUID
	ApplicationID uuid.UUID
	Name          string
	Description   *string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func NewGroup(applicationID uuid.UUID, name string, description *string) *Group {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &Group{
		ID:            newId,
		ApplicationID: applicationID,
		Name:          name,
		Description:   description,
		CreatedAt:     time.Now(),
	}
}
