package removeapplication

import (
	"context"

	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	ApplicationID  uuid.UUID `json:"applicationId" validate:"required"`
	OrganizationID uuid.UUID `json:"organizationId" validate:"required"`
}

type RemoveApplicationService struct {
	ApplicationRepository repository_interfaces.IApplicationRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &RemoveApplicationService{
		ApplicationRepository: repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *RemoveApplicationService) Handler(ctx context.Context, request Request) error {
	err := s.ApplicationRepository.RemoveApplication(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	return nil
}
