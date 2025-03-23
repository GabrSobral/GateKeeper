package login

import (
	"context"

	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	mailservice "github.com/gate-keeper/internal/infra/mail-service"
	"github.com/google/uuid"
)

type Request struct {
	ApplicationID       uuid.UUID `json:"applicationId" validate:"required"`
	Email               string    `json:"email" validate:"required,email"`
	Password            string    `json:"password" validate:"required"`
	CodeChallenge       string    `json:"codeChallenge" validate:"required"`
	CodeChallengeMethod string    `json:"codeChallengeMethod" validate:"required"`
	RedirectUri         string    `json:"redirectUri" validate:"required"`
	ResponseType        string    `json:"responseType" validate:"required"`
	Scope               string    `json:"scope"`
	State               string    `json:"state" validate:"required"`
}

type Response struct {
	MfaEmailRequired   bool    `json:"mfaEmailRequired"`
	MfaAuthAppRequired bool    `json:"mfaAuthAppRequired"`
	SessionCode        *string `json:"sessionCode"`
	Message            string  `json:"message"`
}

type LoginService struct {
	ApplicationUserRepository   repository_interfaces.IApplicationUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	RefreshTokenRepository      repository_interfaces.IRefreshTokenRepository
	AuthozationCodeRepository   repository_interfaces.IApplicationAuthorizationCodeRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository
	EmailMfaCodeRepository      repository_interfaces.IEmailMfaCodeRepository
	SessionCodeRepository       repository_interfaces.ISessionCodeRepository

	MailService mailservice.IMailService
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &LoginService{
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository:    repository_handlers.RefreshTokenRepository{Store: q},
		AuthozationCodeRepository: repository_handlers.ApplicationAuthorizationCodeRepository{Store: q},
		EmailMfaCodeRepository:    repository_handlers.EmailMfaCodeRepository{Store: q},
		SessionCodeRepository:     repository_handlers.SessionCodeRepository{Store: q},

		MailService: &mailservice.MailService{},
	}
}

func (ss *LoginService) Handler(ctx context.Context, request Request) (*Response, error) {
	user, err := ss.ApplicationUserRepository.GetUserByEmail(ctx, request.Email, request.ApplicationID)

	if err != nil {
		return nil, &errors.ErrUserNotFound
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	if !user.IsActive {
		return nil, &errors.ErrUserNotActive
	}

	if user.PasswordHash == nil {
		return nil, &errors.ErrUserSignUpWithSocial
	}

	isPasswordCorrect, err := application_utils.ComparePassword(*user.PasswordHash, request.Password)

	if err != nil {
		return nil, err
	}

	if !isPasswordCorrect {
		return nil, &errors.ErrEmailOrPasswordInvalid
	}

	if !user.IsEmailConfirmed {
		return nil, &errors.ErrEmailNotConfirmed
	}

	if user.IsMfaEmailEnabled {
		userProfile, err := ss.UserProfileRepository.GetUserById(ctx, user.ID)

		if err != nil {
			panic(err)
		}

		emailMfaCode := entities.NewEmailMfaCode(user.ID, user.Email)

		if err := ss.EmailMfaCodeRepository.Add(ctx, emailMfaCode); err != nil {
			panic(err)
		}

		go func() {
			if err := ss.MailService.SendMfaEmail(ctx, user.Email, userProfile.FirstName, emailMfaCode.Token); err != nil {
				panic(err)
			}
		}()

		return &Response{
			MfaEmailRequired:   false,
			MfaAuthAppRequired: true,
			Message:            "MFA is required, please enter the code from your authentication app",
			SessionCode:        nil,
		}, nil
	}

	sessionToken, err := entities.CreateSessionCode(
		user.ID,
		request.ApplicationID,
	)

	if err != nil {
		return nil, err
	}

	if err := ss.SessionCodeRepository.Add(ctx, sessionToken); err != nil {
		return nil, err
	}

	tokenString := sessionToken.Token

	return &Response{
		MfaEmailRequired:   false,
		MfaAuthAppRequired: false,
		Message:            "Login successful",
		SessionCode:        &tokenString,
	}, nil
}
