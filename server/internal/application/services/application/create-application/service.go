package createapplication

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	Name               string    `json:"name" validate:"required,min=3,max=100"`
	Description        *string   `json:"description" validate:"omitempty,min=3,max=100"`
	PasswordHashSecret string    `json:"passwordHashSecret" validate:"required,min=32,max=258"`
	Badges             []string  `json:"badges" validate:"required"`
	HasMfaEmail        bool      `json:"hasMfaEmail" validate:"boolean"`
	HasMfaAuthApp      bool      `json:"hasMfaAuthApp" validate:"boolean"`
	OrganizationID     uuid.UUID `json:"organizationId" validate:"required"`
}

type Response struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	Description        *string   `json:"description"`
	PasswordHashSecret string    `json:"passwordHashSecret"`
	Badges             []string  `json:"badges"`
	HasMfaEmail        bool      `json:"hasMfaEmail"`
	HasMfaAuthApp      bool      `json:"hasMfaAuthApp"`
	OrganizationID     uuid.UUID `json:"organizationId"`
}

type CreateApplicationService struct {
	ApplicationRepository repository_interfaces.IApplicationRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &CreateApplicationService{
		ApplicationRepository: repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *CreateApplicationService) Handler(ctx context.Context, request Request) (*Response, error) {
	newApplication := entities.NewApplication(request.Name, request.Description, request.OrganizationID, request.PasswordHashSecret)

	err := s.ApplicationRepository.AddApplication(ctx, newApplication)

	if err != nil {
		return nil, err
	}

	return &Response{
		ID:                 newApplication.ID,
		Name:               newApplication.Name,
		Description:        newApplication.Description,
		OrganizationID:     newApplication.OrganizationID,
		PasswordHashSecret: newApplication.PasswordHashSecret,
		Badges:             newApplication.Badges,
		HasMfaEmail:        newApplication.HasMfaEmail,
		HasMfaAuthApp:      newApplication.HasMfaAuthApp,
	}, nil
}
