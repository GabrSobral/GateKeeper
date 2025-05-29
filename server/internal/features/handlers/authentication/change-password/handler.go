package changepassword

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/errors"
	application_utils "github.com/gate-keeper/internal/features/utils"
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

func (s *Handler) Handler(ctx context.Context, command Command) error {
	changePasswordCode, err := s.repository.GetChangePasswordCodeByToken(ctx, command.UserID, command.ChangePasswordCode)

	if err != nil {
		return err
	}

	if changePasswordCode == nil {
		return &errors.ErrChangePasswordCodeNotFound
	}

	if changePasswordCode.ExpiresAt.Before(time.Now().UTC()) {
		return &errors.ErrChangePasswordCodeExpired
	}

	if changePasswordCode.Token != command.ChangePasswordCode {
		return &errors.ErrChangePasswordTokenMismatch
	}

	user, err := s.repository.GetUserByID(ctx, changePasswordCode.UserID)

	if err != nil {
		return err
	}

	if user == nil {
		return &errors.ErrUserNotFound
	}

	if user.ShouldChangePass == false {
		return &errors.ErrUserShouldNotChangePassword
	}

	application, err := s.repository.GetApplicationByID(ctx, command.ApplicationID)

	if err != nil {
		return err
	}

	if application == nil {
		return &errors.ErrApplicationNotFound
	}

	hashedPassword, err := application_utils.HashPassword(command.NewPassword, application.PasswordHashSecret)

	if err != nil {
		return err
	}

	user.PasswordHash = &hashedPassword
	user.ShouldChangePass = false

	s.repository.UpdateUser(ctx, user)
	s.repository.RevokeRefreshTokenFromUser(ctx, user.ID)
	s.repository.RevokeAllChangePasswordCodeByUserID(ctx, user.ID)

	return nil
}
