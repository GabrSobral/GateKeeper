package getapplicationuserbyid

import "github.com/google/uuid"

type Query struct {
	UserID uuid.UUID `json:"userId"` // This is the user ID
}
