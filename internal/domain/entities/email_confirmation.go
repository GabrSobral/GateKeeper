package entities

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type EmailConfirmation struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Email     string
	Token     string
	CreatedAt time.Time
	CoolDown  time.Time
	ExpiresAt time.Time
	IsUsed    bool
}

func NewEmailConfirmation(userID uuid.UUID, email string, expiresAt time.Time) *EmailConfirmation {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().UTC()

	return &EmailConfirmation{
		ID:        newId,
		UserID:    userID,
		Email:     email,
		Token:     generateRandomToken(6),
		CoolDown:  currentTime.Add(5 * time.Minute),
		CreatedAt: currentTime,
		ExpiresAt: expiresAt,
		IsUsed:    false,
	}
}

func generateRandomToken(length int16) string {
	const charset = "0123456789"

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
