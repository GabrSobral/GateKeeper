package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type ITenantRepository interface {
	AddTenant(ctx context.Context, tenant *entities.Tenant) error
	RemoveTenant(ctx context.Context, tenantID uuid.UUID) error
	UpdateTenant(ctx context.Context, tenant *entities.Tenant) error
	ListTenants(ctx context.Context) (*[]entities.Tenant, error)
}
