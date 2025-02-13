package listorganizations

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type Response struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListOrganizationsService struct {
	OrganizationRepository repository_interfaces.IOrganizationRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *[]Response] {
	return &ListOrganizationsService{
		OrganizationRepository: repository_handlers.OrganizationRepository{Store: q},
	}
}

func (s *ListOrganizationsService) Handler(ctx context.Context, request Request) (*[]Response, error) {
	organizations := make([]Response, 0)
	organizationsList, err := s.OrganizationRepository.ListOrganizations(ctx)

	if err != nil {
		return nil, err
	}

	for _, organization := range *organizationsList {
		organizations = append(organizations, Response{
			ID:        organization.ID,
			Name:      organization.Name,
			CreatedAt: organization.CreatedAt,
		})
	}

	return &organizations, nil
}
