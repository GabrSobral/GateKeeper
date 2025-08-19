package configureoauthprovider

import (
	"github.com/google/uuid"
)

type Command struct {
	ApplicationID uuid.UUID
	Name          string
	ClientID      string
	ClientSecret  string
	RedirectURI   string
	Enabled       bool
}

type RequestBody struct {
	Name         string `json:"name" validate:"required"`
	ClientID     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	RedirectURI  string `json:"redirectUri" validate:"required,url"`
	Enabled      bool   `json:"enabled"`
}
