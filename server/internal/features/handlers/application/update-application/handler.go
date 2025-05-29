package updateapplication

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
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
	now := time.Now().UTC()

	application := entities.Application{
		ID:                command.ID,
		OrganizationID:    command.OrganizationID,
		Name:              command.Name,
		Description:       command.Description,
		IsActive:          command.IsActive,
		HasMfaAuthApp:     command.HasMfaAuthApp,
		HasMfaEmail:       command.HasMfaEmail,
		Badges:            command.Badges,
		CreatedAt:         now,
		UpdatedAt:         &now,
		CanSelfSignUp:     command.CanSelfSignUp,
		CanSelfForgotPass: command.CanSelfForgotPass,
	}

	err := s.ApplicationRepository.UpdateApplication(ctx, &application)

	if err != nil {
		return err
	}

	return nil
}
