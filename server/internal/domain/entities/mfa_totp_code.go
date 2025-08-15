package entities

import (
	"time"

	"github.com/google/uuid"
)

type MfaTotpCode struct {
	ID          uuid.UUID
	MfaMethodID uuid.UUID
	Secret      string
	CreatedAt   time.Time
}

func NewMfaTotpCode(mfaMethodID uuid.UUID, secret string) *MfaTotpCode {
	newId, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return &MfaTotpCode{
		ID:          newId,
		MfaMethodID: mfaMethodID,
		Secret:      secret,
		CreatedAt:   time.Now().UTC(),
	}
}
