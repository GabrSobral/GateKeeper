package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type ExternalLoginRepository struct {
	Store *pgstore.Queries
}

/*
type IExternalLoginRepository interface {
	GetByProviderKey(ctx context.Context, provider, providerKey string) (*entities.ExternalLogin, error)
	AddExternalLogin(ctx context.Context, externalLogin *entities.ExternalLogin) error
}
*/

func (r ExternalLoginRepository) GetByProviderKey(ctx context.Context, provider, providerKey string) (*entities.ExternalLogin, error) {
	externalLogin, err := r.Store.GetExternalLoginByProviderKey(ctx, pgstore.GetExternalLoginByProviderKeyParams{
		Provider:    provider,
		ProviderKey: providerKey,
	})

	if err != nil {
		return nil, err
	}

	return &entities.ExternalLogin{
		UserID:      externalLogin.UserID,
		Email:       externalLogin.Email,
		Provider:    externalLogin.Provider,
		ProviderKey: externalLogin.ProviderKey,
	}, nil
}

func (r ExternalLoginRepository) AddExternalLogin(ctx context.Context, externalLogin *entities.ExternalLogin) error {
	err := r.Store.AddExternalLogin(ctx, pgstore.AddExternalLoginParams{
		UserID:      externalLogin.UserID,
		Email:       externalLogin.Email,
		Provider:    externalLogin.Provider,
		ProviderKey: externalLogin.ProviderKey,
	})

	return err
}
