package entities

import (
	"time"

	"github.com/google/uuid"
)

type MfaUserSecret struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Secret      string
	IsValidated bool
	CreatedAt   time.Time
	ExpiresAt   time.Time
}

func NewMfaUserSecret(userID uuid.UUID, secret string) *MfaUserSecret {
	newId := uuid.New()

	return &MfaUserSecret{
		ID:          newId,
		UserID:      userID,
		CreatedAt:   time.Now().UTC(),
		Secret:      secret,
		IsValidated: false,
		ExpiresAt:   time.Now().UTC().Add(5 * time.Minute), // Set expiration time to 5 minutes from now
	}
}

func (m *MfaUserSecret) Validate() {
	m.IsValidated = true
	m.ExpiresAt = time.Now().UTC() // Set expiration time to now
}
