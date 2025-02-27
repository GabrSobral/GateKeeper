package getapplicationbyid

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	ApplicationID  uuid.UUID `json:"applicationId" validate:"required"`
	OrganizationID uuid.UUID `json:"organizationId" validate:"required"`
}

type Response struct {
	ID                    uuid.UUID                                  `json:"id"`
	Name                  string                                     `json:"name"`
	Description           *string                                    `json:"description"`
	Badges                []string                                   `json:"badges"`
	CreatedAt             time.Time                                  `json:"createdAt"`
	UpdatedAt             *time.Time                                 `json:"updatedAt"`
	IsActive              bool                                       `json:"isActive"`
	MfaAuthAppEnabled     bool                                       `json:"mfaAuthAppEnabled"`
	MfaEmailEnabled       bool                                       `json:"mfaEmailEnabled"`
	PasswordHashingSecret string                                     `json:"passwordHashingSecret"`
	Secrets               []ApplicationSecrets                       `json:"secrets"`
	Users                 repository_interfaces.ApplicationUsersData `json:"users"`
	Roles                 repository_interfaces.ApplicationRolesData `json:"roles"`
	OAuthProviders        []ApplicationProviders                     `json:"oauthProviders"`
}

type ApplicationSecrets struct {
	ID             uuid.UUID  `json:"id"`
	Name           string     `json:"name"`
	Value          string     `json:"value"`
	ExpirationDate *time.Time `json:"expirationDate"`
}

type ApplicationProviders struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	ClientID     string     `json:"clientId"`
	ClientSecret string     `json:"clientSecret"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	CreatedAt    time.Time  `json:"createdAt"`
}

type GetApplicationByIDService struct {
	OrganizationRepository       repository_interfaces.IOrganizationRepository
	ApplicationRepository        repository_interfaces.IApplicationRepository
	ApplicationUsersRepository   repository_interfaces.IApplicationUserRepository
	ApplicationRolesRepository   repository_interfaces.IApplicationRoleRepository
	ApplicationSecretsRepository repository_interfaces.IApplicationSecretRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &GetApplicationByIDService{
		OrganizationRepository:       repository_handlers.OrganizationRepository{Store: q},
		ApplicationRepository:        repository_handlers.ApplicationRepository{Store: q},
		ApplicationRolesRepository:   repository_handlers.ApplicationRoleRepository{Store: q},
		ApplicationSecretsRepository: repository_handlers.ApplicationSecretRepository{Store: q},
		ApplicationUsersRepository:   repository_handlers.ApplicationUserRepository{Store: q},
	}
}

func (s *GetApplicationByIDService) Handler(ctx context.Context, request Request) (*Response, error) {
	organization, err := s.OrganizationRepository.GetOrganizationByID(ctx, request.OrganizationID)

	if err != nil {
		return nil, err
	}

	if organization == nil {
		return nil, &errors.ErrOrganizationNotFound
	}

	application, err := s.ApplicationRepository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if application == nil {
		return nil, &errors.ErrApplicationNotFound
	}

	secrets := make([]ApplicationSecrets, 0)
	roles := make([]repository_interfaces.ApplicationRoles, 0)

	applicationRolesDb, err := s.ApplicationRolesRepository.ListRolesFromApplication(ctx, application.ID)

	if err != nil {
		return nil, err
	}

	for _, role := range *applicationRolesDb {
		roles = append(roles, repository_interfaces.ApplicationRoles{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	applicationSecretsDb, err := s.ApplicationSecretsRepository.ListSecretsFromApplication(ctx, application.ID)

	if err != nil {
		return nil, err
	}

	for _, secret := range *applicationSecretsDb {
		secrets = append(secrets, ApplicationSecrets{
			ID:             secret.ID,
			Name:           secret.Name,
			Value:          secret.Value[:len(secret.Value)/2] + "****************",
			ExpirationDate: secret.ExpiresAt,
		})
	}

	applicationUsersDb, err := s.ApplicationUsersRepository.GetUsersByApplicationID(ctx, application.ID, 50, 0)

	if err != nil {
		return nil, err
	}

	return &Response{
		ID:                    application.ID,
		Name:                  application.Name,
		Description:           application.Description,
		Badges:                application.Badges,
		CreatedAt:             application.CreatedAt,
		UpdatedAt:             application.UpdatedAt,
		IsActive:              application.IsActive, // to do later
		MfaAuthAppEnabled:     application.HasMfaAuthApp,
		MfaEmailEnabled:       application.HasMfaEmail,
		PasswordHashingSecret: application.PasswordHashSecret,
		Secrets:               secrets,
		Users:                 *applicationUsersDb,
		Roles: repository_interfaces.ApplicationRolesData{
			TotalCount: len(roles),
			Data:       roles,
		},
		OAuthProviders: make([]ApplicationProviders, 0),
	}, nil
}
