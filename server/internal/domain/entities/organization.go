package entities

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID          uuid.UUID
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewOrganization(name string, description *string) *Organization {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &Organization{
		ID:          newId,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now().UTC(),
	}
}
