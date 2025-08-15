package entities

import "github.com/google/uuid"

type MfaWebauthnCredentials struct {
	ID           uuid.UUID
	MfaMethodID  uuid.UUID
	CredentialID string // The credential ID from WebAuthn
	PublicKey    string // The public key associated with the credential
	SignCount    uint32 // The signature counter for the credential
	CreatedAt    int64  // Timestamp when the credential was created
}
