package entities

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationSecret struct {
	ID            uuid.UUID
	ApplicationID uuid.UUID
	Name          string
	Value         string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	ExpiresAt     *time.Time
}

func NewApplicationSecret(applicationID uuid.UUID, name string, expiresAt *time.Time) *ApplicationSecret {
	newID, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &ApplicationSecret{
		ID:            newID,
		ApplicationID: applicationID,
		Name:          name,
		Value:         GenerateRandomString(32),
		CreatedAt:     time.Now(),
		UpdatedAt:     nil,
		ExpiresAt:     expiresAt,
	}
}
