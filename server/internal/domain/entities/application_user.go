package entities

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationUser struct {
	ID               uuid.UUID
	ApplicationID    uuid.UUID
	Email            string
	PasswordHash     *string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	IsActive         bool
	IsEmailConfirmed bool
	TwoFactorEnabled bool
	TwoFactorSecret  *string
}

func CreateUser(email string, passwordHash *string, applicationID uuid.UUID) (*ApplicationUser, error) {
	userId, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &ApplicationUser{
		ID:               userId,
		ApplicationID:    applicationID,
		Email:            email,
		PasswordHash:     passwordHash,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        nil,
		IsActive:         true,
		IsEmailConfirmed: false,
		TwoFactorEnabled: false,
		TwoFactorSecret:  nil,
	}, nil
}

func NewUser(applicationID, id uuid.UUID, email string, passwordHash *string, createdAt time.Time, updatedAt *time.Time, isActive, isEmailConfirmed, twoFactorEnabled bool, twoFactorSecret *string) *ApplicationUser {
	return &ApplicationUser{
		ID:               id,
		ApplicationID:    applicationID,
		Email:            email,
		PasswordHash:     passwordHash,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		IsActive:         isActive,
		IsEmailConfirmed: isEmailConfirmed,
		TwoFactorEnabled: twoFactorEnabled,
		TwoFactorSecret:  twoFactorSecret,
	}
}
