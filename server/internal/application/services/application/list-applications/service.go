package listapplications

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
	OrganizationID uuid.UUID `json:"organizationId" validate:"required"`
}

type Response struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Badges      []string  `json:"badges"`
}

type ListApplicationsService struct {
	OrganizationRepository repository_interfaces.IOrganizationRepository
	ApplicationRepository  repository_interfaces.IApplicationRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *[]Response] {
	return &ListApplicationsService{
		OrganizationRepository: repository_handlers.OrganizationRepository{Store: q},
		ApplicationRepository:  repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *ListApplicationsService) Handler(ctx context.Context, request Request) (*[]Response, error) {
	organization, err := s.OrganizationRepository.GetOrganizationByID(ctx, request.OrganizationID)

	if err != nil {
		return nil, err
	}

	if organization == nil {
		return nil, &errors.ErrOrganizationNotFound
	}

	applications, err := s.ApplicationRepository.ListApplicationsFromOrganization(ctx, organization.ID)

	if err != nil {
		return nil, err
	}

	response := make([]Response, 0)

	for _, application := range *applications {
		if len(application.Badges) == 1 && application.Badges[0] == "" {
			application.Badges = make([]string, 0)
		}

		response = append(response, Response{
			ID:          application.ID,
			Name:        application.Name,
			Description: application.Description,
			Badges:      application.Badges,
		})
	}

	return &response, nil
}
