package repository_handlers

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	Store *pgstore.Queries
}

func (r UserRepository) AddUser(ctx context.Context, newUser *entities.User) error {
	err := r.Store.AddUser(ctx, pgstore.AddUserParams{
		ID:               newUser.ID,
		Email:            newUser.Email,
		PasswordHash:     newUser.PasswordHash,
		CreatedAt:        pgtype.Timestamp{Time: newUser.CreatedAt, Valid: true},
		UpdatedAt:        newUser.UpdatedAt,
		IsActive:         newUser.IsActive,
		IsEmailConfirmed: newUser.IsEmailConfirmed,
		TwoFactorEnabled: newUser.TwoFactorEnabled,
		TwoFactorSecret:  newUser.TwoFactorSecret,
	})

	return err
}

type GetUserByEmailParams struct {
	Email string
}

func (r UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := r.Store.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:               user.ID,
		Email:            user.Email,
		PasswordHash:     user.PasswordHash,
		CreatedAt:        user.CreatedAt.Time,
		UpdatedAt:        user.UpdatedAt,
		IsActive:         user.IsActive,
		IsEmailConfirmed: user.IsEmailConfirmed,
		TwoFactorEnabled: user.TwoFactorEnabled,
		TwoFactorSecret:  user.TwoFactorSecret,
	}, nil
}

func (r UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, err := r.Store.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:               user.ID,
		Email:            user.Email,
		PasswordHash:     user.PasswordHash,
		CreatedAt:        user.CreatedAt.Time,
		UpdatedAt:        user.UpdatedAt,
		IsActive:         user.IsActive,
		IsEmailConfirmed: user.IsEmailConfirmed,
		TwoFactorEnabled: user.TwoFactorEnabled,
		TwoFactorSecret:  user.TwoFactorSecret,
	}, nil
}

func (r UserRepository) IsUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	_, err := r.Store.GetUserByEmail(ctx, email)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r UserRepository) IsUserExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	_, err := r.Store.GetUserById(ctx, id)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r UserRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	now := time.Now().UTC()

	err := r.Store.UpdateUser(ctx, pgstore.UpdateUserParams{
		ID:               user.ID,
		Email:            user.Email,
		PasswordHash:     user.PasswordHash,
		UpdatedAt:        &now,
		IsActive:         user.IsActive,
		IsEmailConfirmed: user.IsEmailConfirmed,
		TwoFactorEnabled: user.TwoFactorEnabled,
		TwoFactorSecret:  user.TwoFactorSecret,
	})

	return user, err
}
