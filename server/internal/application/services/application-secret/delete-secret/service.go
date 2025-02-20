package deletesecret

import (
	"context"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	SecretID      uuid.UUID `json:"secretId" validate:"required,uuid"`
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
}

type DeleteSecretService struct {
	ApplicationRepository       repository_interfaces.IApplicationRepository
	ApplicationSecretRepository repository_interfaces.IApplicationSecretRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &DeleteSecretService{
		ApplicationSecretRepository: repository_handlers.ApplicationSecretRepository{Store: q},
		ApplicationRepository:       repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *DeleteSecretService) Handler(ctx context.Context, request Request) error {
	isApplicationExists, err := s.ApplicationRepository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	if !isApplicationExists {
		return &errors.ErrAplicationNotFound
	}

	if err := s.ApplicationSecretRepository.RemoveSecret(ctx, request.SecretID); err != nil {
		return err
	}

	return nil
}
