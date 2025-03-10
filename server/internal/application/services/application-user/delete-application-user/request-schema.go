package deleteapplicationuser

import "github.com/google/uuid"

type Request struct {
	ApplicationID uuid.UUID
	UserID        uuid.UUID
}
