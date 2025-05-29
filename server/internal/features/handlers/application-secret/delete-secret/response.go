package deletesecret

import "github.com/google/uuid"

type Response struct {
	ID uuid.UUID `json:"id"`
	// ApplicationID uuid.UUID  `json:"applicationId"`
	// Name          string     `json:"name"`
	// Value         string     `json:"value"`
	// CreatedAt     time.Time  `json:"createdAt"`
	// UpdatedAt     *time.Time `json:"updatedAt"`
	// ExpiresAt     *time.Time `json:"expirationDate"`
}
