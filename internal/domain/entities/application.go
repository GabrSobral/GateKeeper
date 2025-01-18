package entities

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func NewApplication(name string, description *string, tenantID uuid.UUID) *Application {
	newID, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &Application{
		ID:          newID,
		TenantID:    tenantID,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
}
