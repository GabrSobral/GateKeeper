package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TenantRepository struct {
	Store *pgstore.Queries
}

/*
type ITenantRepository interface {
	AddTenant(ctx context.Context, tenant *entities.Tenant) error
	RemoveTenant(ctx context.Context, tenantID string) error
	UpdateTenant(ctx context.Context, tenant *entities.Tenant) error
	ListTenants(ctx context.Context) (*[]entities.Tenant, error)
}
*/

func (r TenantRepository) AddTenant(ctx context.Context, tenant *entities.Tenant) error {
	err := r.Store.AddTenant(ctx, pgstore.AddTenantParams{
		UserID:      tenant.ID,
		Name:        tenant.Name,
		Description: tenant.Description,
		CreatedAt:   pgtype.Timestamp{Time: tenant.CreatedAt, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r TenantRepository) RemoveTenant(ctx context.Context, tenantID uuid.UUID) error {
	err := r.Store.RemoveTenant(ctx, tenantID)
	if err != nil {
		return err
	}

	return nil
}

func (r TenantRepository) UpdateTenant(ctx context.Context, tenant *entities.Tenant) error {
	err := r.Store.UpdateTenant(ctx, pgstore.UpdateTenantParams{
		ID:          tenant.ID,
		Name:        tenant.Name,
		UpdatedAt:   tenant.UpdatedAt,
		Description: tenant.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r TenantRepository) ListTenants(ctx context.Context) (*[]entities.Tenant, error) {
	tenants, err := r.Store.ListTenants(ctx)

	if err != nil {
		return nil, err
	}

	var tenantList []entities.Tenant
	for _, tenant := range tenants {
		tenantList = append(tenantList, entities.Tenant{
			ID:          tenant.ID,
			Name:        tenant.Name,
			CreatedAt:   tenant.CreatedAt.Time,
			UpdatedAt:   tenant.UpdatedAt,
			Description: tenant.Description,
		})
	}

	return &tenantList, nil
}
