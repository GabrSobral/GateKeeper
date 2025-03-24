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

type CreateApplicationService struct {
	ApplicationRepository     repository_interfaces.IApplicationRepository
	ApplicationRoleRepository repository_interfaces.IApplicationRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &CreateApplicationService{
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},
		ApplicationRoleRepository: repository_handlers.ApplicationRoleRepository{Store: q},
	}
}

func (s *CreateApplicationService) Handler(ctx context.Context, request Request) (*Response, error) {
	newApplication := entities.AddApplication(
		request.Name,
		request.Description,
		request.OrganizationID,
		request.PasswordHashSecret,
		request.Badges,
		request.HasMfaEmail,
		request.HasMfaAuthApp,
		true, // IsActive
		nil,  // UpdatedAt
		request.CanSelfSignUp,
		request.CanSelfForgotPass,
	)

	err := s.ApplicationRepository.AddApplication(ctx, newApplication)

	if err != nil {
		return nil, err
	}

	userRoleDescription := "Default user role"
	adminRoleDescription := "Default admin role"

	userRole := entities.NewApplicationRole(newApplication.ID, "User", &userRoleDescription)
	adminRole := entities.NewApplicationRole(newApplication.ID, "Admin", &adminRoleDescription)

	// Add default roles
	s.ApplicationRoleRepository.AddRole(ctx, userRole)
	s.ApplicationRoleRepository.AddRole(ctx, adminRole)

	return &Response{
		ID:                 newApplication.ID,
		Name:               newApplication.Name,
		Description:        newApplication.Description,
		OrganizationID:     newApplication.OrganizationID,
		PasswordHashSecret: newApplication.PasswordHashSecret,
		Badges:             newApplication.Badges,
		HasMfaEmail:        newApplication.HasMfaEmail,
		HasMfaAuthApp:      newApplication.HasMfaAuthApp,
		IsActive:           newApplication.IsActive,
		CanSelfSignUp:      newApplication.CanSelfSignUp,
		CanSelfForgotPass:  newApplication.CanSelfForgotPass,
	}, nil
}
