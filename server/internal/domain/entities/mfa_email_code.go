package entities

import (
	"time"

	"github.com/google/uuid"
)

type MfaEmailCode struct {
	ID          uuid.UUID
	MfaMethodID uuid.UUID
	Token       string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	Verified    bool
}

func NewMfaEmailCode(mfaMethodID uuid.UUID) *MfaEmailCode {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().UTC()
	expiresAt := currentTime.Add(15 * time.Minute) // 15 minutes

	return &MfaEmailCode{
		ID:          newId,
		MfaMethodID: mfaMethodID,
		Token:       generateRandomToken(6),
		CreatedAt:   currentTime,
		ExpiresAt:   expiresAt,
		Verified:    false,
	}
}
