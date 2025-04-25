package verifymfa

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	Code          string    `json:"code" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	ApplicationID uuid.UUID `json:"applicationId" validate:"required"`
}

type Response struct {
	SessionCode string `json:"sessionCode"`
}

type VerifyMfaService struct {
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	UserProfileRepository     repository_interfaces.IUserProfileRepository
	RefreshTokenRepository    repository_interfaces.IRefreshTokenRepository
	AuthozationCodeRepository repository_interfaces.IApplicationAuthorizationCodeRepository
	EmailMfaCodeRepository    repository_interfaces.IEmailMfaCodeRepository
	SessionCodeRepository     repository_interfaces.ISessionCodeRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &VerifyMfaService{
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository:    repository_handlers.RefreshTokenRepository{Store: q},
		AuthozationCodeRepository: repository_handlers.ApplicationAuthorizationCodeRepository{Store: q},
		EmailMfaCodeRepository:    repository_handlers.EmailMfaCodeRepository{Store: q},
		SessionCodeRepository:     repository_handlers.SessionCodeRepository{Store: q},
	}
}

func (ss *VerifyMfaService) Handler(ctx context.Context, request Request) (*Response, error) {
	user, err := ss.ApplicationUserRepository.GetUserByEmail(ctx, request.Email, request.ApplicationID)

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
		return nil, &errors.ErrMfaEmailNotEnabled
	}

	emailMfaCode, err := ss.EmailMfaCodeRepository.GetByToken(ctx, user.ID, request.Code)

	if err != nil {
		return nil, &errors.ErrEmailMfaCodeNotFound
	}

	if emailMfaCode == nil {
		return nil, &errors.ErrEmailMfaCodeNotFound
	}

	if emailMfaCode.ExpiresAt.Before(time.Now().UTC()) {
		return nil, &errors.ErrEmailMfaCodeExpired
	}

	ss.EmailMfaCodeRepository.DeleteByID(ctx, emailMfaCode.ID)

	sessionCode, err := entities.CreateSessionCode(user.ID, request.ApplicationID)

	if err := ss.SessionCodeRepository.Add(ctx, sessionCode); err != nil {
		return nil, err
	}

	return &Response{
		SessionCode: sessionCode.Token,
	}, nil
}
