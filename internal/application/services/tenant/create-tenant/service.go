package createtenant

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

type CreateTenantService struct {
	TenantRepository repository_interfaces.ITenantRepository
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
	return &CreateTenantService{
		TenantRepository: repository_handlers.TenantRepository{Store: q},
	}
}

func (s *CreateTenantService) Handler(ctx context.Context, request Request) (*Response, error) {
	newTenant := entities.NewTenant(request.Name, request.Description)

	if err := s.TenantRepository.AddTenant(ctx, newTenant); err != nil {
		return nil, err
	}

	return &Response{
		ID:          newTenant.ID,
		Name:        newTenant.Name,
		Description: newTenant.Description,
		CreatedAt:   newTenant.CreatedAt,
		UpdatedAt:   newTenant.UpdatedAt,
	}, nil
}
