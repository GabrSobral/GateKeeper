package repository_handlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
	pgstore "github.com/guard-service/internal/infra/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	lop "github.com/samber/lo"
)

type UserRepository struct {
	Store *pgstore.Queries
}

func (r UserRepository) AddUser(ctx context.Context, newUser *entities.User) error {

	err := r.Store.AddUser(ctx, pgstore.AddUserParams{
		ID:               newUser.ID,
		Email:            newUser.Email,
		PasswordHash:     pgtype.Text{String: *newUser.PasswordHash, Valid: true},
		CreatedAt:        pgtype.Timestamp{Time: newUser.CreatedAt, Valid: true},
		UpdatedAt:        pgtype.Timestamp{Time: *newUser.UpdatedAt, Valid: true},
		IsActive:         newUser.IsActive,
		IsEmailConfirmed: newUser.IsEmailConfirmed,
		TwoFactorEnabled: newUser.TwoFactorEnabled,
		TwoFactorSecret:  pgtype.Text{String: *newUser.TwoFactorSecret, Valid: true},
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
		PasswordHash:     lop.Ternary(user.PasswordHash.Valid, &user.PasswordHash.String, nil),
		CreatedAt:        user.CreatedAt.Time,
		UpdatedAt:        lop.Ternary(user.UpdatedAt.Valid, &user.UpdatedAt.Time, nil),
		IsActive:         user.IsActive,
		IsEmailConfirmed: user.IsEmailConfirmed,
		TwoFactorEnabled: user.TwoFactorEnabled,
		TwoFactorSecret:  lop.Ternary(user.TwoFactorSecret.Valid, &user.TwoFactorSecret.String, nil),
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
		PasswordHash:     lop.Ternary(user.PasswordHash.Valid, &user.PasswordHash.String, nil),
		CreatedAt:        user.CreatedAt.Time,
		UpdatedAt:        lop.Ternary(user.UpdatedAt.Valid, &user.UpdatedAt.Time, nil),
		IsActive:         user.IsActive,
		IsEmailConfirmed: user.IsEmailConfirmed,
		TwoFactorEnabled: user.TwoFactorEnabled,
		TwoFactorSecret:  lop.Ternary(user.TwoFactorSecret.Valid, &user.TwoFactorSecret.String, nil),
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
	err := r.Store.UpdateUser(ctx, pgstore.UpdateUserParams{
		ID:               user.ID,
		Email:            user.Email,
		PasswordHash:     lop.Ternary(user.PasswordHash != nil, pgtype.Text{String: *user.PasswordHash, Valid: true}, pgtype.Text{}),
		UpdatedAt:        pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		IsActive:         user.IsActive,
		IsEmailConfirmed: user.IsEmailConfirmed,
		TwoFactorEnabled: user.TwoFactorEnabled,
		TwoFactorSecret:  pgtype.Text{String: *user.TwoFactorSecret, Valid: true},
	})

	return user, err
}
