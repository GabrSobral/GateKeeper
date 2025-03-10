package repository_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationUserRepository struct {
	Store *pgstore.Queries
}

func (r ApplicationUserRepository) DeleteApplicationUser(ctx context.Context, applicationID, userID uuid.UUID) error {
	err := r.Store.DeleteApplicationUser(ctx, pgstore.DeleteApplicationUserParams{
		ID:            userID,
		ApplicationID: applicationID,
	})

	return err
}

func (r ApplicationUserRepository) GetUsersByApplicationID(ctx context.Context, applicationID uuid.UUID, limit, offset int) (*repository_interfaces.ApplicationUsersData, error) {
	users, err := r.Store.GetUsersByApplicationID(ctx, pgstore.GetUsersByApplicationIDParams{
		ApplicationID: applicationID,
		Limit:         int32(limit),
		Offset:        int32(offset),
	})

	if err != nil {
		return nil, err
	}

	totalUsers := 0

	if len(users) > 0 {
		totalUsers = int(users[0].TotalUsers)
	}

	result := repository_interfaces.ApplicationUsersData{
		TotalCount: totalUsers,
		Data:       []repository_interfaces.ApplicationUsers{},
	}

	for _, user := range users {
		roles := []repository_interfaces.ApplicationRoles{}

		fmt.Println(string(user.Roles))

		err := json.Unmarshal(user.Roles, &roles)

		if err != nil {
			return nil, err
		}

		result.Data = append(result.Data, repository_interfaces.ApplicationUsers{
			ID:          user.ID,
			DisplayName: *user.DisplayName,
			Email:       user.Email,
			Roles:       roles,
		})
	}

	return &result, nil
}

func (r ApplicationUserRepository) AddUser(ctx context.Context, newUser *entities.ApplicationUser) error {
	err := r.Store.AddUser(ctx, pgstore.AddUserParams{
		ID:                  newUser.ID,
		Email:               newUser.Email,
		ApplicationID:       newUser.ApplicationID,
		ShouldChangePass:    newUser.ShouldChangePass,
		PasswordHash:        newUser.PasswordHash,
		CreatedAt:           pgtype.Timestamp{Time: newUser.CreatedAt, Valid: true},
		UpdatedAt:           newUser.UpdatedAt,
		IsActive:            newUser.IsActive,
		IsEmailConfirmed:    newUser.IsEmailConfirmed,
		IsMfaAuthAppEnabled: newUser.IsMfaAuthAppEnabled,
		IsMfaEmailEnabled:   newUser.IsMfaEmailEnabled,
		TwoFactorSecret:     newUser.TwoFactorSecret,
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
		ID:                  user.ID,
		Email:               user.Email,
		PasswordHash:        user.PasswordHash,
		CreatedAt:           user.CreatedAt.Time,
		UpdatedAt:           user.UpdatedAt,
		IsActive:            user.IsActive,
		IsEmailConfirmed:    user.IsEmailConfirmed,
		IsMfaAuthAppEnabled: user.IsMfaAuthAppEnabled,
		ApplicationID:       user.ApplicationID,
		ShouldChangePass:    user.ShouldChangePass,
		IsMfaEmailEnabled:   user.IsMfaEmailEnabled,
		TwoFactorSecret:     user.TwoFactorSecret,
	}, nil
}

func (r ApplicationUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.ApplicationUser, error) {
	user, err := r.Store.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.ApplicationUser{
		ID:                  user.ID,
		Email:               user.Email,
		PasswordHash:        user.PasswordHash,
		CreatedAt:           user.CreatedAt.Time,
		UpdatedAt:           user.UpdatedAt,
		IsActive:            user.IsActive,
		IsEmailConfirmed:    user.IsEmailConfirmed,
		IsMfaAuthAppEnabled: user.IsMfaAuthAppEnabled,
		IsMfaEmailEnabled:   user.IsMfaEmailEnabled,
		ApplicationID:       user.ApplicationID,
		ShouldChangePass:    user.ShouldChangePass,
		TwoFactorSecret:     user.TwoFactorSecret,
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
		ID:                  user.ID,
		Email:               user.Email,
		PasswordHash:        user.PasswordHash,
		UpdatedAt:           &now,
		IsActive:            user.IsActive,
		IsEmailConfirmed:    user.IsEmailConfirmed,
		IsMfaAuthAppEnabled: user.IsMfaAuthAppEnabled,
		IsMfaEmailEnabled:   user.IsMfaEmailEnabled,
		TwoFactorSecret:     user.TwoFactorSecret,
		ShouldChangePass:    user.ShouldChangePass,
	})

	return user, err
}
