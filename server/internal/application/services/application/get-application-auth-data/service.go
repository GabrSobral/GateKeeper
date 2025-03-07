package getapplicationauthdata

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
	ApplicationID uuid.UUID `json:"applicationId" validate:"required"`
}

type Response struct {
	ID                uuid.UUID              `json:"id"`
	Name              string                 `json:"name"`
	CanSelfSignUp     bool                   `json:"canSelfSignUp"`
	CanSelfForgotPass bool                   `json:"canSelfForgotPass"`
	OAuthProviders    []ApplicationProviders `json:"oauthProviders"`
}

type ApplicationProviders struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	ClientID     string    `json:"clientId"`
	ClientSecret string    `json:"clientSecret"`
}

type GetApplicationAuthDataService struct {
	OrganizationRepository repository_interfaces.IOrganizationRepository
	ApplicationRepository  repository_interfaces.IApplicationRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &GetApplicationAuthDataService{
		OrganizationRepository: repository_handlers.OrganizationRepository{Store: q},
		ApplicationRepository:  repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *GetApplicationAuthDataService) Handler(ctx context.Context, request Request) (*Response, error) {
	application, err := s.ApplicationRepository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if application == nil {
		return nil, &errors.ErrApplicationNotFound
	}

	return &Response{
		ID:                application.ID,
		Name:              application.Name,
		CanSelfSignUp:     application.CanSelfSignUp,
		CanSelfForgotPass: application.CanSelfForgotPass,
		OAuthProviders:    make([]ApplicationProviders, 0),
	}, nil
}
