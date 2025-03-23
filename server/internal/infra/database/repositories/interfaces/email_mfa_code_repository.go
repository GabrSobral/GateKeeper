package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IEmailMfaCodeRepository interface {
	Add(ctx context.Context, emailMfaCode *entities.EmailMfaCode) error
	GetByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.EmailMfaCode, error)
	Update(ctx context.Context, emailMfaCode *entities.EmailMfaCode) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
