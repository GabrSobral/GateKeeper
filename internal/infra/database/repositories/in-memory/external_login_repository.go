package inmemory_repositories

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
)

type InMemoryExternalLoginRepository struct {
	Logins map[string]*entities.ExternalLogin
}

func (r InMemoryExternalLoginRepository) GetByProviderKey(ctx context.Context, provider, providerKey string) (*entities.ExternalLogin, error) {
	for _, login := range r.Logins {
		if login.Provider == provider && login.ProviderKey == providerKey {
			return login, nil
		}
	}

	return nil, nil
}

func (r InMemoryExternalLoginRepository) AddExternalLogin(ctx context.Context, externalLogin *entities.ExternalLogin) error {
	r.Logins[externalLogin.UserID.String()] = externalLogin

	return nil
}
