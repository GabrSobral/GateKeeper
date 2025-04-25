package deleteapplicationuser

import (
	"context"

	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error)
	DeleteApplicationUser(ctx context.Context, applicationID, userID uuid.UUID) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error) {
	isApplicationExists, err := r.Store.CheckIfApplicationExists(ctx, applicationID)

	if err != nil {
		return false, err
	}

	return isApplicationExists, nil
}

func (r Repository) DeleteApplicationUser(ctx context.Context, applicationID, userID uuid.UUID) error {
	err := r.Store.DeleteApplicationUser(ctx, pgstore.DeleteApplicationUserParams{
		ID:            userID,
		ApplicationID: applicationID,
	})

	return err
}
