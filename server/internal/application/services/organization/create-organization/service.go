package createorganization

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type CreateOrganizationService struct {
	OrganizationRepository repository_interfaces.IOrganizationRepository
}

type Request struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description" validate:"omitempty"`
}

type Response struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &CreateOrganizationService{
		OrganizationRepository: repository_handlers.OrganizationRepository{Store: q},
	}
}

func (s *CreateOrganizationService) Handler(ctx context.Context, request Request) (*Response, error) {
	newOrganization := entities.NewOrganization(request.Name, request.Description)

	if err := s.OrganizationRepository.AddOrganization(ctx, newOrganization); err != nil {
		return nil, err
	}

	return &Response{
		ID:          newOrganization.ID,
		Name:        newOrganization.Name,
		Description: newOrganization.Description,
		CreatedAt:   newOrganization.CreatedAt,
		UpdatedAt:   newOrganization.UpdatedAt,
	}, nil
}
