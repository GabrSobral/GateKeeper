package externalloginprovider

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	UserID           uuid.UUID  `json:"user_id"`
	UserEmail        string     `json:"user_email"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	CreatedAt        time.Time  `json:"created_at"`
	IsEmailConfirmed bool       `json:"is_email_confirmed"`
	UpdatedAt        *time.Time `json:"updated_at"`
}
