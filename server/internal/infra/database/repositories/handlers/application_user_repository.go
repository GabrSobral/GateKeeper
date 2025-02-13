package repository_handlers

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationUserRepository struct {
	Store *pgstore.Queries
}

func (r ApplicationUserRepository) AddUser(ctx context.Context, newUser *entities.ApplicationUser) error {
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

func (r ApplicationUserRepository) GetUserByEmail(ctx context.Context, email string, applicationID uuid.UUID) (*entities.ApplicationUser, error) {
	user, err := r.Store.GetUserByEmail(ctx, pgstore.GetUserByEmailParams{
		Email:         email,
		ApplicationID: applicationID,
	})

	if err != nil {
		return nil, err
	}

	return &entities.ApplicationUser{
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

func (r ApplicationUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.ApplicationUser, error) {
	user, err := r.Store.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.ApplicationUser{
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

func (r ApplicationUserRepository) IsUserExistsByEmail(ctx context.Context, email string, applicationID uuid.UUID) (bool, error) {
	_, err := r.Store.GetUserByEmail(ctx, pgstore.GetUserByEmailParams{
		Email:         email,
		ApplicationID: applicationID,
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r ApplicationUserRepository) IsUserExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	_, err := r.Store.GetUserById(ctx, id)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r ApplicationUserRepository) UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error) {
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
