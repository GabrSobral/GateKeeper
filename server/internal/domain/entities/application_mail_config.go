package entities

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationMailConfig struct {
	ApplicationID uuid.UUID
	Host          string
	Port          int
	Username      string
	Password      string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func NewApplicationMailConfig(applicationID uuid.UUID, host, username, password string, port int) *ApplicationMailConfig {
	return &ApplicationMailConfig{
		ApplicationID: applicationID,
		Host:          host,
		Port:          port,
		Username:      username,
		Password:      password,
		CreatedAt:     time.Now(),
		UpdatedAt:     nil,
	}
}
