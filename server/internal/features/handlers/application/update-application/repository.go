package updateapplication

import (
	"context"
	"strings"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type IRepository interface {
	UpdateApplication(ctx context.Context, application *entities.Application) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) UpdateApplication(ctx context.Context, newApplication *entities.Application) error {
	badges := strings.Join(newApplication.Badges, ",")

	err := r.Store.UpdateApplication(ctx, pgstore.UpdateApplicationParams{
		ID:                newApplication.ID,
		Name:              newApplication.Name,
		Description:       newApplication.Description,
		HasMfaAuthApp:     newApplication.HasMfaAuthApp,
		Badges:            &badges,
		IsActive:          newApplication.IsActive,
		HasMfaEmail:       newApplication.HasMfaEmail,
		UpdatedAt:         newApplication.UpdatedAt,
		CanSelfSignUp:     newApplication.CanSelfSignUp,
		CanSelfForgotPass: newApplication.CanSelfForgotPass,
	})

	if err != nil {
		return err
	}

	return nil
}
