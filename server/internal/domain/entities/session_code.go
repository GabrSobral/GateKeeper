package entities

import (
	"time"

	"github.com/google/uuid"
)

type SessionCode struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Token         string
	ExpiresAt     time.Time
	CreatedAt     time.Time
	ApplicationID uuid.UUID
	IsUsed        bool
}

func CreateSessionCode(userID, applicationID uuid.UUID) (*SessionCode, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	currentTime := time.Now().UTC()
	expiresAt := currentTime.Add(15 * time.Minute) // 15 minutes

	return &SessionCode{
		ID:            id,
		UserID:        userID,
		ApplicationID: applicationID,
		ExpiresAt:     expiresAt,
		CreatedAt:     currentTime,
		Token:         GenerateRandomString(56),
		IsUsed:        false,
	}, nil
}
