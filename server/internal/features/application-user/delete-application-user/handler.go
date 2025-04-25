package deleteapplicationuser

import (
	"context"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type Handler struct {
	repository IRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Command] {
	return &Handler{
		repository: Repository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, request Command) error {
	isApplicationExists, err := s.repository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	if !isApplicationExists {
		return &errors.ErrApplicationNotFound
	}

	if err := s.repository.DeleteApplicationUser(ctx, request.ApplicationID, request.UserID); err != nil {
		return err
	}

	return nil
}
