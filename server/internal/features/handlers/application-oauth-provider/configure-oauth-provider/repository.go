package configureoauthprovider

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	AddApplicationOauthProvider(ctx context.Context, applicationOauthProvider *entities.ApplicationOAuthProvider) error
	UpdateApplicationOauthProvider(ctx context.Context, applicationOauthProvider *entities.ApplicationOAuthProvider) error
	CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error)
	GetApplicationOauthProviderByName(ctx context.Context, applicationID uuid.UUID, name string) (*entities.ApplicationOAuthProvider, error)
	CheckIfApplicationOauthProviderConfigurationExists(ctx context.Context, applicationID uuid.UUID, name string) (bool, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) AddApplicationOauthProvider(ctx context.Context, applicationOauthProvider *entities.ApplicationOAuthProvider) error {
	err := r.Store.AddApplicationOauthProvider(ctx, pgstore.AddApplicationOauthProviderParams{
		ID:            applicationOauthProvider.ID,
		ApplicationID: applicationOauthProvider.ApplicationID,
		Name:          applicationOauthProvider.Name,
		ClientID:      applicationOauthProvider.ClientID,
		ClientSecret:  applicationOauthProvider.ClientSecret,
		RedirectUri:   applicationOauthProvider.RedirectURI,
		CreatedAt:     pgtype.Timestamp{Time: applicationOauthProvider.CreatedAt, Valid: true},
		UpdatedAt:     applicationOauthProvider.UpdatedAt,
		Enabled:       applicationOauthProvider.Enabled,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpdateApplicationOauthProvider(ctx context.Context, applicationOauthProvider *entities.ApplicationOAuthProvider) error {
	err := r.Store.UpdateApplicationOauthProvider(ctx, pgstore.UpdateApplicationOauthProviderParams{
		ID:           applicationOauthProvider.ID,
		Name:         applicationOauthProvider.Name,
		ClientID:     applicationOauthProvider.ClientID,
		ClientSecret: applicationOauthProvider.ClientSecret,
		RedirectUri:  applicationOauthProvider.RedirectURI,
		UpdatedAt:    applicationOauthProvider.UpdatedAt,
		Enabled:      applicationOauthProvider.Enabled,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error) {
	exists, err := r.Store.CheckIfApplicationExists(ctx, applicationID)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r Repository) GetApplicationOauthProviderByName(ctx context.Context, applicationID uuid.UUID, name string) (*entities.ApplicationOAuthProvider, error) {
	applicationOauthProvider, err := r.Store.GetApplicationOauthProviderByName(ctx, pgstore.GetApplicationOauthProviderByNameParams{
		ApplicationID: applicationID,
		Name:          name,
	})

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.ApplicationOAuthProvider{
		ID:            applicationOauthProvider.ID,
		ApplicationID: applicationOauthProvider.ApplicationID,
		Name:          applicationOauthProvider.Name,
		ClientID:      applicationOauthProvider.ClientID,
		ClientSecret:  applicationOauthProvider.ClientSecret,
		RedirectURI:   applicationOauthProvider.RedirectUri,
		CreatedAt:     applicationOauthProvider.CreatedAt.Time,
		UpdatedAt:     applicationOauthProvider.UpdatedAt,
		Enabled:       applicationOauthProvider.Enabled,
	}, nil
}

func (r Repository) CheckIfApplicationOauthProviderConfigurationExists(ctx context.Context, applicationID uuid.UUID, name string) (bool, error) {
	exists, err := r.Store.CheckIfApplicationOauthProviderConfigurationExists(ctx, pgstore.CheckIfApplicationOauthProviderConfigurationExistsParams{
		ApplicationID: applicationID,
		Name:          name,
	})

	if err != nil {
		return false, err
	}

	return exists, nil
}
