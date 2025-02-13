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
	ID                    uuid.UUID            `json:"id"`
	Name                  string               `json:"name"`
	Description           *string              `json:"description"`
	Badges                []string             `json:"badges"`
	CreatedAt             time.Time            `json:"createdAt"`
	UpdatedAt             *time.Time           `json:"updatedAt"`
	DeactivatedAt         *time.Time           `json:"deactivatedAt"`
	MfaAuthAppEnabled     bool                 `json:"mfaAuthAppEnabled"`
	MfaEmailEnabled       bool                 `json:"mfaEmailEnabled"`
	PasswordHashingSecret string               `json:"passwordHashingSecret"`
	Secrets               []applicationSecrets `json:"secrets"`
	Users                 applicationUsersData `json:"users"`
	Roles                 applicationRolesData `json:"roles"`
}

type applicationSecrets struct {
	ID             uuid.UUID  `json:"id"`
	Name           string     `json:"name"`
	Value          string     `json:"value"`
	ExpirationDate *time.Time `json:"expirationDate"`
}

type applicationUsersData struct {
	TotalCount int                `json:"totalCount"`
	Data       []applicationUsers `json:"data"`
}

type applicationRolesData struct {
	TotalCount int                `json:"totalCount"`
	Data       []applicationRoles `json:"data"`
}

type applicationUsers struct {
	ID          uuid.UUID          `json:"id"`
	DisplayName string             `json:"displayName"`
	Email       string             `json:"email"`
	Roles       []applicationRoles `json:"roles"`
}

type applicationRoles struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

type GetApplicationByIDService struct {
	OrganizationRepository repository_interfaces.IOrganizationRepository
	ApplicationRepository  repository_interfaces.IApplicationRepository
	ApplicationUsers       repository_interfaces.IApplicationUserRepository
	ApplicationRoles       repository_interfaces.IApplicationRoleRepository
	ApplicationSecrets     repository_interfaces.IApplicationSecretRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &GetApplicationByIDService{
		OrganizationRepository: repository_handlers.OrganizationRepository{Store: q},
		ApplicationRepository:  repository_handlers.ApplicationRepository{Store: q},
		ApplicationRoles:       repository_handlers.ApplicationRoleRepository{Store: q},
		ApplicationSecrets:     repository_handlers.ApplicationSecretRepository{Store: q},
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
		return nil, &errors.ErrAplicationNotFound
	}

	secrets := make([]applicationSecrets, 0)
	users := make([]applicationUsers, 0)
	roles := make([]applicationRoles, 0)

	applicationRolesDb, err := s.ApplicationRoles.ListRolesFromApplication(ctx, application.ID)

	if err != nil {
		return nil, err
	}

	for _, role := range *applicationRolesDb {
		roles = append(roles, applicationRoles{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	applicationSecretsDb, err := s.ApplicationSecrets.ListSecretsFromApplication(ctx, application.ID)

	if err != nil {
		return nil, err
	}

	for _, secret := range *applicationSecretsDb {
		secrets = append(secrets, applicationSecrets{
			ID:             secret.ID,
			Name:           secret.Name,
			Value:          secret.Value[:len(secret.Value)/2] + "****************",
			ExpirationDate: secret.ExpiresAt,
		})
	}

	return &Response{
		ID:                    application.ID,
		Name:                  application.Name,
		Description:           application.Description,
		Badges:                application.Badges,
		CreatedAt:             application.CreatedAt,
		UpdatedAt:             application.UpdatedAt,
		DeactivatedAt:         nil, // to do later
		MfaAuthAppEnabled:     application.HasMfaAuthApp,
		MfaEmailEnabled:       application.HasMfaEmail,
		PasswordHashingSecret: application.PasswordHashSecret,
		Secrets:               secrets,
		Users: applicationUsersData{
			TotalCount: len(users),
			Data:       users,
		},
		Roles: applicationRolesData{
			TotalCount: len(roles),
			Data:       roles,
		},
	}, nil
}
