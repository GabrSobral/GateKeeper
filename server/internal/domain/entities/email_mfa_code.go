package entities

import (
	"time"

	"github.com/google/uuid"
)

type EmailMfaCode struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Email     string
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
	IsUsed    bool
}

func NewEmailMfaCode(userID uuid.UUID, email string) *EmailMfaCode {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().UTC()
	expiresAt := currentTime.Add(15 * time.Minute) // 15 minutes

	return &EmailMfaCode{
		ID:        newId,
		UserID:    userID,
		Email:     email,
		Token:     generateRandomToken(6),
		CreatedAt: currentTime,
		ExpiresAt: expiresAt,
		IsUsed:    false,
	}
}
