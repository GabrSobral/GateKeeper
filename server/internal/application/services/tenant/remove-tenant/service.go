package removetenant

import (
	"context"

	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type RemoveTenantService struct {
	TenantRepository repository_interfaces.ITenantRepository
}

type Request struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &RemoveTenantService{
		TenantRepository: repository_handlers.TenantRepository{Store: q},
	}
}

func (s *RemoveTenantService) Handler(ctx context.Context, request Request) error {
	if err := s.TenantRepository.RemoveTenant(ctx, request.ID); err != nil {
		return err
	}

	return nil
}
