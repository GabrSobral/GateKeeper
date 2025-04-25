package deleteapplicationuser

import (
	"github.com/google/uuid"
)

type Command struct {
	ApplicationID uuid.UUID
	UserID        uuid.UUID
}
