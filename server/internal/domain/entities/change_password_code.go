package entities

import (
	"time"

	"github.com/google/uuid"
)

type ChangePasswordCode struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Email     string
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewChangePasswordCode(userID uuid.UUID, email string) *ChangePasswordCode {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().UTC()
	expiresAt := currentTime.Add(15 * time.Minute) // 15 minutes

	return &ChangePasswordCode{
		ID:        newId,
		UserID:    userID,
		Email:     email,
		Token:     GenerateRandomString(64),
		CreatedAt: currentTime,
		ExpiresAt: expiresAt,
	}
}
