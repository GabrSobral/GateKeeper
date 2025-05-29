package verifyappmfa

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	mailservice "github.com/gate-keeper/internal/infra/mail-service"
	"github.com/pquerna/otp/totp"
)

type Handler struct {
	repository  IRepository
	mailService mailservice.IMailService
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Command, *Response] {
	return &Handler{
		repository:  Repository{Store: q},
		mailService: &mailservice.MailService{},
	}
}

func (s *Handler) Handler(ctx context.Context, command Command) (*Response, error) {
	user, err := s.repository.GetUserByEmail(ctx, command.Email, command.ApplicationID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	if !user.IsActive {
		return nil, &errors.ErrUserNotActive
	}

	if !user.IsEmailConfirmed {
		return nil, &errors.ErrEmailNotConfirmed
	}

	if !user.IsMfaEmailEnabled {
		return nil, &errors.ErrMfaAppNotEnabled
	}

	if user.TwoFactorSecret == nil {
		return nil, &errors.ErrMfaAppNotEnabled
	}

	appMfaCode, err := s.repository.GetAppMfaCodeByID(ctx, *command.MfaID)

	if err != nil {
		return nil, &errors.ErrAppMfaCodeNotFound
	}

	if appMfaCode == nil {
		return nil, &errors.ErrAppMfaCodeNotFound
	}

	if appMfaCode.ExpiresAt.Before(time.Now().UTC()) {
		return nil, &errors.ErrAppMfaCodeExpired
	}

	if err := s.repository.DeleteAppMfaCodeByID(ctx, appMfaCode.ID); err != nil {
		return nil, err
	}

	isValid := totp.Validate(command.Code, *user.TwoFactorSecret)

	if !isValid {
		return nil, &errors.ErrInvalidMfaAuthAppCode
	}

	sessionCode, err := entities.CreateSessionCode(user.ID, command.ApplicationID)

	if err != nil {
		return nil, err
	}

	if err := s.repository.AddSessionCode(ctx, sessionCode); err != nil {
		return nil, err
	}

	return &Response{
		SessionCode: sessionCode.Token,
	}, nil
}
