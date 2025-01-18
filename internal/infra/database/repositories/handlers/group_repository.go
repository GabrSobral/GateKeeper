package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type GroupRepository struct {
	Store *pgstore.Queries
}

func (r GroupRepository) AddGroup(ctx context.Context, group *entities.Group) error {
	err := r.Store.AddGroup(ctx, pgstore.AddGroupParams{
		ID:            group.ID,
		Name:          group.Name,
		Description:   group.Description,
		ApplicationID: group.ApplicationID,
		CreatedAt:     pgtype.Timestamp{Time: group.CreatedAt, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r GroupRepository) GetGroupByID(ctx context.Context, groupID uuid.UUID) (*entities.Group, error) {
	group, err := r.Store.GetGroupById(ctx, groupID)

	if err != nil {
		return nil, err
	}

	return &entities.Group{
		ID:            group.ID,
		Name:          group.Name,
		Description:   group.Description,
		ApplicationID: group.ApplicationID,
		CreatedAt:     group.CreatedAt.Time,
		UpdatedAt:     group.UpdatedAt,
	}, nil
}

func (r GroupRepository) RemoveGroup(ctx context.Context, groupID uuid.UUID) error {
	err := r.Store.RemoveGroup(ctx, groupID)

	if err != nil {
		return err
	}

	return nil
}

func (r GroupRepository) UpdateGroup(ctx context.Context, group *entities.Group) error {
	err := r.Store.UpdateGroup(ctx, pgstore.UpdateGroupParams{
		ID:            group.ID,
		Name:          group.Name,
		ApplicationID: group.ApplicationID,
		Description:   group.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r GroupRepository) ListGroupsFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.Group, error) {
	groups, err := r.Store.ListGroupsFromApplication(ctx, applicationID)

	if err != nil {
		return nil, err
	}

	var groupsList []entities.Group

	for _, group := range groups {
		groupsList = append(groupsList, entities.Group{
			ID:            group.ID,
			Name:          group.Name,
			Description:   group.Description,
			ApplicationID: group.ApplicationID,
		})
	}

	return &groupsList, nil
}
