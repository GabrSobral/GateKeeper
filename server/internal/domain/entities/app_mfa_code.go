package entities

import (
	"time"

	"github.com/google/uuid"
)

type AppMfaCode struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Email     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewAppMfaCode(userID uuid.UUID, email string) *AppMfaCode {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().UTC()
	expiresAt := currentTime.Add(15 * time.Minute) // 15 minutes

	return &AppMfaCode{
		ID:        newId,
		UserID:    userID,
		Email:     email,
		CreatedAt: currentTime,
		ExpiresAt: expiresAt,
	}
}
