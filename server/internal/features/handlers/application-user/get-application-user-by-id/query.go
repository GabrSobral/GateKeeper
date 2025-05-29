package getapplicationuserbyid

import "github.com/google/uuid"

type Query struct {
	UserID uuid.UUID
}
