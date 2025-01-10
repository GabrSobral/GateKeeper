package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID
	Email            string
	PasswordHash     *string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	IsActive         bool
	IsEmailConfirmed bool
	TwoFactorEnabled bool
	TwoFactorSecret  *string
}

func CreateUser(email string, passwordHash *string) (*User, error) {
	userId, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &User{
		ID:               userId,
		Email:            email,
		PasswordHash:     passwordHash,
		CreatedAt:        time.Now().UTC(),
		IsActive:         true,
		IsEmailConfirmed: false,
		TwoFactorEnabled: false,
		TwoFactorSecret:  nil,
	}, nil
}

func NewUser(id uuid.UUID, email string, passwordHash *string, createdAt time.Time, updatedAt *time.Time, isActive, isEmailConfirmed, twoFactorEnabled bool, twoFactorSecret *string) *User {
	return &User{
		ID:               id,
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
