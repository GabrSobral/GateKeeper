package signupcredential

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken uuid.UUID    `json:"refreshToken"`
}

type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	DisplayName   string    `json:"displayName"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	ApplicationID uuid.UUID `json:"applicationId"`
	Email         string    `json:"email"`
	PhotoURL      *string   `json:"photoUrl"`
	CreatedAt     time.Time `json:"createdAt"`
}
