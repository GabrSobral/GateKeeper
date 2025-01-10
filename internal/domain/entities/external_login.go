package entities

import "github.com/google/uuid"

type ExternalLogin struct {
	UserID      uuid.UUID
	Email       string
	Provider    string
	ProviderKey string
}

func NewExternalLogin(userID uuid.UUID, userEmail, provider, providerKey string) *ExternalLogin {
	return &ExternalLogin{
		UserID:      userID,
		Email:       userEmail,
		Provider:    provider,
		ProviderKey: providerKey,
	}
}
