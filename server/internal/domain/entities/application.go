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

func NewApplication(ID uuid.UUID, name string, description *string, organizationID uuid.UUID, passwordHashSecret string, badges []string, hasMfaEmail, hasMfaAuthApp, isActive bool, updatedAt *time.Time, createdAt time.Time) *Application {
	return &Application{
		ID:                 ID,
		OrganizationID:     organizationID,
		Name:               name,
		Description:        description,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
		PasswordHashSecret: passwordHashSecret,
		IsActive:           isActive,
		HasMfaAuthApp:      hasMfaAuthApp,
		HasMfaEmail:        hasMfaEmail,
		Badges:             badges,
	}
}

func AddApplication(name string, description *string, organizationID uuid.UUID, passwordHashSecret string, badges []string, hasMfaEmail, hasMfaAuthApp, isActive bool, updatedAt *time.Time) *Application {
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
		UpdatedAt:          updatedAt,
		PasswordHashSecret: passwordHashSecret,
		IsActive:           isActive,
		HasMfaAuthApp:      hasMfaAuthApp,
		HasMfaEmail:        hasMfaEmail,
		Badges:             badges,
	}
}
