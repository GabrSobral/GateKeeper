package confirmmfaauthappsecret

import (
	"context"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/pquerna/otp/totp"
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
	user, err := s.repository.GetUserByID(ctx, command.UserID)

	if err != nil {
		return err
	}

	if user == nil {
		return &errors.ErrUserNotFound
	}

	if !user.IsMfaAuthAppEnabled {
		return &errors.ErrMfaAuthAppNotEnabled
	}

	mfaUserSecret, err := s.repository.GetMfaUserSecretByUserID(ctx, user.ID)

	if err != nil {
		return err
	}

	if mfaUserSecret == nil {
		return &errors.ErrMfaUserSecretNotFound
	}

	// if mfaUserSecret.IsValidated {
	// 	return &errors.ErrMfaUserSecretAlreadyValidated
	// }

	isValid := totp.Validate(command.MfaAuthAppCode, mfaUserSecret.Secret)

	if !isValid {
		return &errors.ErrInvalidMfaAuthAppCode
	}

	// mfaUserSecret.Validate()

	// if err := s.repository.UpdateMfaUserSecret(ctx, mfaUserSecret); err != nil {
	// 	return err
	// }

	if err := s.repository.RevokeMfaUserSecret(ctx, command.UserID); err != nil {
		return err
	}

	user.TwoFactorSecret = &mfaUserSecret.Secret
	s.repository.UpdateUser(ctx, user)

	return nil
}
