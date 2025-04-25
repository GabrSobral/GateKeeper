package createapplication

import "github.com/google/uuid"

type Response struct {
	ID uuid.UUID `json:"id"`
}
