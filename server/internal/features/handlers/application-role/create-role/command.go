package createrole

import "github.com/google/uuid"

type Command struct {
	ApplicationID uuid.UUID
	Name          string
	Description   *string
}

type RequestBody struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}
