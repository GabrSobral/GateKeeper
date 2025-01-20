package creategroup

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type CreateGroupService struct {
	GroupRepository repository_interfaces.IGroupRepository
}

type Request struct {
	Name          string    `json:"name" validate:"required"`
	Description   *string   `json:"description" validate:"omitempty"`
	ApplicationID uuid.UUID `json:"application_id" validate:"required"`
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &CreateGroupService{
		GroupRepository: repository_handlers.GroupRepository{Store: q},
	}
}

func (s *CreateGroupService) Handler(ctx context.Context, request Request) error {
	group := entities.NewGroup(request.ApplicationID, request.Name, request.Description)

	if err := s.GroupRepository.AddGroup(ctx, group); err != nil {
		return err
	}

	return nil
}
