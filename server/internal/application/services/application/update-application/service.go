package updateapplication

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

type Request struct {
	ID                 uuid.UUID `json:"id" validate:"required"`
	Name               string    `json:"name" validate:"required,min=3,max=100"`
	Description        *string   `json:"description" validate:"omitempty,min=3,max=100"`
	PasswordHashSecret string    `json:"passwordHashSecret" validate:"required,min=32,max=258"`
	Badges             []string  `json:"badges" validate:"required"`
	HasMfaEmail        bool      `json:"hasMfaEmail" validate:"boolean"`
	HasMfaAuthApp      bool      `json:"hasMfaAuthApp" validate:"boolean"`
	OrganizationID     uuid.UUID `json:"organizationId" validate:"required"`
	IsActive           bool      `json:"isActive" validate:"required"`
	CanSelfSignUp      bool      `json:"canSelfSignUp" validate:"boolean"`
	CanSelfForgotPass  bool      `json:"canSelfForgotPass" validate:"boolean"`
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
	IsActive           bool      `json:"isActive"`
	CanSelfSignUp      bool      `json:"canSelfSignUp"`
	CanSelfForgotPass  bool      `json:"canSelfForgotPass"`
}

type UpdateApplicationService struct {
	ApplicationRepository repository_interfaces.IApplicationRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &UpdateApplicationService{
		ApplicationRepository: repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *UpdateApplicationService) Handler(ctx context.Context, request Request) (*Response, error) {
	now := time.Now()

	application := entities.Application{
		ID:                 request.ID,
		OrganizationID:     request.OrganizationID,
		Name:               request.Name,
		Description:        request.Description,
		IsActive:           request.IsActive,
		HasMfaAuthApp:      request.HasMfaAuthApp,
		HasMfaEmail:        request.HasMfaEmail,
		PasswordHashSecret: request.PasswordHashSecret,
		Badges:             request.Badges,
		CreatedAt:          now,
		UpdatedAt:          &now,
		CanSelfSignUp:      request.CanSelfSignUp,
		CanSelfForgotPass:  request.CanSelfForgotPass,
	}

	err := s.ApplicationRepository.UpdateApplication(ctx, &application)

	if err != nil {
		return nil, err
	}

	return &Response{
		ID:                 application.ID,
		Name:               application.Name,
		Description:        application.Description,
		OrganizationID:     application.OrganizationID,
		PasswordHashSecret: application.PasswordHashSecret,
		Badges:             application.Badges,
		HasMfaEmail:        application.HasMfaEmail,
		HasMfaAuthApp:      application.HasMfaAuthApp,
		IsActive:           application.IsActive,
		CanSelfSignUp:      application.CanSelfSignUp,
		CanSelfForgotPass:  application.CanSelfForgotPass,
	}, nil
}
