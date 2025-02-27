package deleterole

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
	RoleID        uuid.UUID `json:"roleId" validate:"required,uuid"`
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
}

type DeleteRoleService struct {
	ApplicationRepository     repository_interfaces.IApplicationRepository
	ApplicationRoleRepository repository_interfaces.IApplicationRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &DeleteRoleService{
		ApplicationRoleRepository: repository_handlers.ApplicationRoleRepository{Store: q},
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *DeleteRoleService) Handler(ctx context.Context, request Request) error {
	isApplicationExists, err := s.ApplicationRepository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	if !isApplicationExists {
		return &errors.ErrApplicationNotFound
	}

	if err := s.ApplicationRoleRepository.RemoveRole(ctx, request.RoleID); err != nil {
		return err
	}

	return nil
}
