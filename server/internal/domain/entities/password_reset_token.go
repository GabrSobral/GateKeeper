package entities

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type PasswordResetToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewPasswordResetToken(userID uuid.UUID) (*PasswordResetToken, error) {
	newID, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &PasswordResetToken{
		ID:        newID,
		UserID:    userID,
		Token:     GenerateRandomString(128),
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(time.Minute * 15),
	}, nil
}

func GenerateRandomString(length int16) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
