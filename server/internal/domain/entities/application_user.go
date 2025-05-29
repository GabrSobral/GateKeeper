package entities

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationUser struct {
	ID                  uuid.UUID
	ApplicationID       uuid.UUID
	Email               string
	PasswordHash        *string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
	IsActive            bool
	IsEmailConfirmed    bool
	ShouldChangePass    bool
	IsMfaAuthAppEnabled bool
	IsMfaEmailEnabled   bool
	TwoFactorSecret     *string
	Preferred2FAMethod  *int
}

func CreateApplicationUser(email string, passwordHash *string, applicationID uuid.UUID, shouldChangePass bool) (*ApplicationUser, error) {
	userId, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &ApplicationUser{
		ID:                  userId,
		ApplicationID:       applicationID,
		Email:               email,
		PasswordHash:        passwordHash,
		CreatedAt:           time.Now().UTC(),
		UpdatedAt:           nil,
		IsActive:            true,
		ShouldChangePass:    shouldChangePass,
		IsEmailConfirmed:    false,
		IsMfaAuthAppEnabled: false,
		IsMfaEmailEnabled:   false,
		TwoFactorSecret:     nil,
		Preferred2FAMethod:  nil,
	}, nil
}

func NewApplicationUser(applicationID, id uuid.UUID, email string, passwordHash *string, createdAt time.Time, updatedAt *time.Time, isActive, isEmailConfirmed, IsMfaEmailEnabled, IsMfaAuthAppEnabled bool, twoFactorSecret *string, shouldChangePass bool) *ApplicationUser {
	return &ApplicationUser{
		ID:                  id,
		ApplicationID:       applicationID,
		Email:               email,
		PasswordHash:        passwordHash,
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		IsActive:            isActive,
		IsEmailConfirmed:    isEmailConfirmed,
		IsMfaAuthAppEnabled: IsMfaAuthAppEnabled,
		IsMfaEmailEnabled:   IsMfaEmailEnabled,
		TwoFactorSecret:     twoFactorSecret,
		ShouldChangePass:    shouldChangePass,
	}
}
