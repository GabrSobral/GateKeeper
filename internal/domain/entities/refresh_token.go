package entities

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID                 uuid.UUID
	UserID             uuid.UUID
	AvailableRefreshes uint8
	ExpiresAt          time.Time
	CreatedAt          time.Time
}

func CreateRefreshToken(userID uuid.UUID, availableRefreshes uint8, expiresAt time.Time) (*RefreshToken, error) {
	id, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &RefreshToken{
		ID:                 id,
		UserID:             userID,
		AvailableRefreshes: availableRefreshes,
		ExpiresAt:          expiresAt,
		CreatedAt:          time.Now().UTC(),
	}, nil
}

func NewRefreshToken(ID uuid.UUID, userID uuid.UUID, availableRefreshes uint8, expiresAt time.Time) *RefreshToken {
	return &RefreshToken{
		ID:                 ID,
		UserID:             userID,
		AvailableRefreshes: availableRefreshes,
		ExpiresAt:          expiresAt,
		CreatedAt:          time.Now().UTC(),
	}
}
