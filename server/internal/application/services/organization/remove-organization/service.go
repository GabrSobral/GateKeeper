package removeorganization

import (
	"context"

	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type RemoveOrganizationService struct {
	OrganizationRepository repository_interfaces.IOrganizationRepository
}

type Request struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &RemoveOrganizationService{
		OrganizationRepository: repository_handlers.OrganizationRepository{Store: q},
	}
}

func (s *RemoveOrganizationService) Handler(ctx context.Context, request Request) error {
	if err := s.OrganizationRepository.RemoveOrganization(ctx, request.ID); err != nil {
		return err
	}

	return nil
}
