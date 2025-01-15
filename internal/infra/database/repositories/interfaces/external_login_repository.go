package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
)

type IExternalLoginRepository interface {
	GetByProviderKey(ctx context.Context, provider, providerKey string) (*entities.ExternalLogin, error)
	AddExternalLogin(ctx context.Context, externalLogin *entities.ExternalLogin) error
}
