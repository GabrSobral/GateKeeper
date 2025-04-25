package resetpassword

import (
	"context"
	"time"

	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/errors"
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
	passwordResetToken, err := s.repository.GetPasswordResetByTokenID(ctx, command.PasswordResetId)

	if err != nil {
		return err
	}

	if passwordResetToken == nil {
		return &errors.ErrPasswordResetNotFound
	}

	if passwordResetToken.ExpiresAt.Before(time.Now().UTC()) {
		return &errors.ErrPasswordResetTokenExpired
	}

	if passwordResetToken.Token != command.PasswordResetToken {
		return &errors.ErrPasswordResetTokenMismatch
	}

	user, err := s.repository.GetUserByID(ctx, passwordResetToken.UserID)

	if err != nil {
		return err
	}

	if user == nil {
		return &errors.ErrUserNotFound
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

	s.repository.UpdateUser(ctx, user)
	s.repository.RevokeRefreshTokenFromUser(ctx, user.ID)
	s.repository.DeletePasswordResetFromUser(ctx, user.ID)

	return nil
}
