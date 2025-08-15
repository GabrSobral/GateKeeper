package entities

import (
	"time"

	"github.com/google/uuid"
)

const (
	MfaMethodTotp  = "totp"
	MfaMethodEmail = "email"
	MfaMethodSms   = "sms"
)

type MfaMethod struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Type       string // e.g., "email", "sms", "totp"
	Enabled    bool
	CreatedAt  time.Time
	LastUsedAt *time.Time // Nullable, can be nil if never used
}
