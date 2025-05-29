package listroles

import "github.com/google/uuid"

type Response struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}
