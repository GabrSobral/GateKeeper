package listroles

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
	ApplicationID  uuid.UUID `json:"applicationId" validate:"required,uuid"`
	OrganizationID uuid.UUID `json:"organizationId" validate:"required,uuid"`
}

type Response struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

type ListRolesService struct {
	ApplicationRepository     repository_interfaces.IApplicationRepository
	ApplicationRoleRepository repository_interfaces.IApplicationRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *[]Response] {
	return &ListRolesService{
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},
		ApplicationRoleRepository: repository_handlers.ApplicationRoleRepository{Store: q},
	}
}

func (s *ListRolesService) Handler(ctx context.Context, request Request) (*[]Response, error) {
	isApplicationExists, err := s.ApplicationRepository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if !isApplicationExists {
		return nil, &errors.ErrApplicationNotFound
	}

	roles, err := s.ApplicationRoleRepository.ListRolesFromApplication(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	var response []Response

	for _, role := range *roles {
		response = append(response, Response{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	return &response, nil
}
