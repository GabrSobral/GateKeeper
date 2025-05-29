package getapplicationbyid

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID                    uuid.UUID              `json:"id"`
	Name                  string                 `json:"name"`
	Description           *string                `json:"description"`
	Badges                []string               `json:"badges"`
	CreatedAt             time.Time              `json:"createdAt"`
	UpdatedAt             *time.Time             `json:"updatedAt"`
	IsActive              bool                   `json:"isActive"`
	MfaAuthAppEnabled     bool                   `json:"mfaAuthAppEnabled"`
	MfaEmailEnabled       bool                   `json:"mfaEmailEnabled"`
	CanSelfSignUp         bool                   `json:"canSelfSignUp"`
	CanSelfForgotPass     bool                   `json:"canSelfForgotPass"`
	PasswordHashingSecret string                 `json:"passwordHashingSecret"`
	Secrets               []ApplicationSecrets   `json:"secrets"`
	Users                 ApplicationUsersData   `json:"users"`
	Roles                 ApplicationRolesData   `json:"roles"`
	OAuthProviders        []ApplicationProviders `json:"oauthProviders"`
}

type ApplicationSecrets struct {
	ID             uuid.UUID  `json:"id"`
	Name           string     `json:"name"`
	Value          string     `json:"value"`
	ExpirationDate *time.Time `json:"expirationDate"`
}

type ApplicationProviders struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	ClientID     string     `json:"clientId"`
	ClientSecret string     `json:"clientSecret"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	CreatedAt    time.Time  `json:"createdAt"`
}
