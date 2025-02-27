package deleteapplicationuser

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
	ApplicationID uuid.UUID
	UserID        uuid.UUID
}

type applicationRoles struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

type DeleteApplicationUserService struct {
	ApplicationRepository     repository_interfaces.IApplicationRepository
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	UserRoleRepository        repository_interfaces.IUserRoleRepository
	UserProfileRepository     repository_interfaces.IUserProfileRepository
	ApplicationRoles          repository_interfaces.IApplicationRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &DeleteApplicationUserService{
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserRoleRepository:        repository_handlers.UserRoleRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
		ApplicationRoles:          repository_handlers.ApplicationRoleRepository{Store: q},
	}
}

func (s *DeleteApplicationUserService) Handler(ctx context.Context, request Request) error {
	isApplicationExists, err := s.ApplicationRepository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	if !isApplicationExists {
		return &errors.ErrApplicationNotFound
	}

	if err := s.ApplicationUserRepository.DeleteApplicationUser(ctx, request.ApplicationID, request.UserID); err != nil {
		return err
	}

	return nil
}
