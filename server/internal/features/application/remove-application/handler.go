package removeapplication

import (
	"context"

	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type Handler struct {
	ApplicationRepository repository_interfaces.IApplicationRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Command] {
	return &Handler{
		ApplicationRepository: repository_handlers.ApplicationRepository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, command Command) error {
	err := s.ApplicationRepository.RemoveApplication(ctx, command.ApplicationID)

	if err != nil {
		return err
	}

	return nil
}
