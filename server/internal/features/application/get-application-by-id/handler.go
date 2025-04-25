package getapplicationbyid

import (
	"context"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type Handler struct {
	repository IRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Query, *Response] {
	return &Handler{
		repository: Repository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, request Query) (*Response, error) {
	application, err := s.repository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if application == nil {
		return nil, &errors.ErrApplicationNotFound
	}

	secrets := make([]ApplicationSecrets, 0)
	roles := make([]ApplicationRoles, 0)

	applicationRolesDb, err := s.repository.ListRolesFromApplication(ctx, application.ID)

	if err != nil {
		return nil, err
	}

	for _, role := range *applicationRolesDb {
		roles = append(roles, ApplicationRoles{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	applicationSecretsDb, err := s.repository.ListSecretsFromApplication(ctx, application.ID)

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

	applicationUsersDb, err := s.repository.GetUsersByApplicationID(ctx, application.ID, 50, 0)

	if err != nil {
		return nil, err
	}

	if len(application.Badges) == 1 && application.Badges[0] == "" {
		application.Badges = make([]string, 0)
	}

	return &Response{
		ID:                    application.ID,
		Name:                  application.Name,
		Description:           application.Description,
		Badges:                application.Badges,
		CreatedAt:             application.CreatedAt,
		UpdatedAt:             application.UpdatedAt,
		CanSelfSignUp:         application.CanSelfSignUp,
		CanSelfForgotPass:     application.CanSelfForgotPass,
		IsActive:              application.IsActive,
		MfaAuthAppEnabled:     application.HasMfaAuthApp,
		MfaEmailEnabled:       application.HasMfaEmail,
		PasswordHashingSecret: application.PasswordHashSecret,
		Secrets:               secrets,
		Users:                 *applicationUsersDb,
		Roles: ApplicationRolesData{
			TotalCount: len(roles),
			Data:       roles,
		},
		OAuthProviders: make([]ApplicationProviders, 0),
	}, nil
}
