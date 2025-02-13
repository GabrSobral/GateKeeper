package entities

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationOAuthProvider struct {
	ID           uuid.UUID
	AppicationID uuid.UUID
	Name         string
	ClientID     uuid.UUID
	ClientSecret string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func NewApplicationOAuthProvider(applicationID uuid.UUID, name string, clientID uuid.UUID, clientSecret string) *ApplicationOAuthProvider {
	newID, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &ApplicationOAuthProvider{
		ID:           newID,
		AppicationID: applicationID,
		Name:         name,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		CreatedAt:    time.Now(),
		UpdatedAt:    nil,
	}
}
