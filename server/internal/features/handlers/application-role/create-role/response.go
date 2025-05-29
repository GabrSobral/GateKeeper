package createrole

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Description   *string    `json:"description"`
	ApplicationID uuid.UUID  `json:"applicationId"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}
