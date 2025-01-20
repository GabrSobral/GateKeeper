package removegroup

import (
	"context"

	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type RemoveGroupService struct {
	GroupRepository repository_interfaces.IGroupRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &RemoveGroupService{
		GroupRepository: repository_handlers.GroupRepository{Store: q},
	}
}

type Request struct {
	GroupID uuid.UUID `json:"group_id"`
}

func (s *RemoveGroupService) Handler(ctx context.Context, request Request) error {
	err := s.GroupRepository.RemoveGroup(ctx, request.GroupID)

	if err != nil {
		return err
	}

	return nil
}
