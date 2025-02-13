package createrole

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
	Name          string    `json:"name" validate:"required"`
	Description   *string   `json:"description"`
}

type Response struct {
	ID            uuid.UUID
	ApplicationID uuid.UUID
	Name          string
	Description   *string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

type CreateRoleService struct {
	ApplicationRepository     repository_interfaces.IApplicationRepository
	ApplicationRoleRepository repository_interfaces.IApplicationRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &CreateRoleService{
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},
		ApplicationRoleRepository: repository_handlers.ApplicationRoleRepository{Store: q},
	}
}

func (s *CreateRoleService) Handler(ctx context.Context, request Request) (*Response, error) {
	isApplicationExists, err := s.ApplicationRepository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if !isApplicationExists {
		return nil, &errors.ErrAplicationNotFound
	}

	newRole := entities.NewApplicationRole(request.ApplicationID, request.Name, request.Description)

	if err := s.ApplicationRoleRepository.AddRole(ctx, newRole); err != nil {
		return nil, err
	}

	return &Response{
		ID:            newRole.ID,
		Name:          newRole.Name,
		Description:   newRole.Description,
		ApplicationID: newRole.ApplicationID,
		CreatedAt:     newRole.CreatedAt,
		UpdatedAt:     newRole.UpdatedAt,
	}, nil
}
