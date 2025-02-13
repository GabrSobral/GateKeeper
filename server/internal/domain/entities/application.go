package entities

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	Name               string
	Description        *string
	IsActive           bool
	HasMfaAuthApp      bool
	HasMfaEmail        bool
	PasswordHashSecret string
	Badges             []string
	CreatedAt          time.Time
	UpdatedAt          *time.Time
}

func NewApplication(name string, description *string, organizationID uuid.UUID, passwordHashSecret string) *Application {
	newID, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &Application{
		ID:                 newID,
		OrganizationID:     organizationID,
		Name:               name,
		Description:        description,
		CreatedAt:          time.Now(),
		UpdatedAt:          nil,
		PasswordHashSecret: passwordHashSecret,
		IsActive:           true,
		HasMfaAuthApp:      false,
		HasMfaEmail:        false,
		Badges:             []string{},
	}
}
