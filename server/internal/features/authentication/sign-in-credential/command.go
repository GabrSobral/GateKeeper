package signincredential

import "github.com/google/uuid"

type Command struct {
	AuthorizationCode uuid.UUID `json:"authorizationCode"`
	ClientSecret      string    `json:"clientSecret"`
	ClientID          uuid.UUID `json:"clientId"`
	CodeVerifier      string    `json:"codeVerifier"`
	RedirectURI       string    `json:"redirectUri"`
}
