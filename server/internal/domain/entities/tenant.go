package entities

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID          uuid.UUID
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewTenant(name string, description *string) *Tenant {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &Tenant{
		ID:          newId,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now().UTC(),
	}
}
