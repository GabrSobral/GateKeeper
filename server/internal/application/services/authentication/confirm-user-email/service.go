package confirmuseremail

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
	Token               string    `json:"token" validate:"required"`
	Email               string    `json:"email" validate:"required,email"`
	ApplicationID       uuid.UUID `json:"applicationId" validate:"required"`
	CodeChallengeMethod string    `json:"codeChallengeMethod" validate:"required"`
	ResponseType        string    `json:"responseType" validate:"required"`
	Scope               string    `json:"scope" validate:"required"`
	State               string    `json:"state" validate:"required"`
	CodeChallenge       string    `json:"codeChallenge" validate:"required"`
	RedirectUri         string    `json:"redirectUri" validate:"required"`
}

type Response struct {
	AuthorizationCode   string `json:"authorizationCode"`
	RedirectUri         string `json:"redirectUri"`
	State               string `json:"state"`
	CodeChallenge       string `json:"codeChallenge"`
	CodeChallengeMethod string `json:"codeChallengeMethod"`
	ResponseType        string `json:"responseType"`
}

type ConfirmUserEmail struct {
	ApplicationUserRepository   repository_interfaces.IApplicationUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository
	RefreshTokenRepository      repository_interfaces.IRefreshTokenRepository
	AuthozationCodeRepository   repository_interfaces.IApplicationAuthorizationCodeRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &ConfirmUserEmail{
		ApplicationUserRepository:   repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:       repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository:      repository_handlers.RefreshTokenRepository{Store: q},
		EmailConfirmationRepository: repository_handlers.EmailConfirmationRepository{Store: q},
		AuthozationCodeRepository:   repository_handlers.ApplicationAuthorizationCodeRepository{Store: q},
	}
}

func (cm *ConfirmUserEmail) Handler(ctx context.Context, request Request) (*Response, error) {
	user, err := cm.ApplicationUserRepository.GetUserByEmail(ctx, request.Email, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	authorizationCode, err := entities.CreateApplicationAuthorizationCode(
		request.ApplicationID,
		user.ID,
		request.RedirectUri,
		request.CodeChallenge,
		request.CodeChallengeMethod,
	)

	if err != nil {
		return nil, err
	}

	if err := cm.AuthozationCodeRepository.AddAuthorizationCode(ctx, authorizationCode); err != nil {
		return nil, err
	}

	emailConfirmation, err := cm.EmailConfirmationRepository.GetByEmail(ctx, request.Email, user.ID)

	if err != nil {
		return nil, nil
	}

	if emailConfirmation == nil {
		return nil, &errors.ErrEmailConfirmationNotFound
	}

	if emailConfirmation.Token != request.Token {
		return nil, &errors.ErrConfirmationTokenInvalid
	}

	if emailConfirmation.IsUsed {
		return nil, &errors.ErrConfirmationTokenAlreadyUsed
	}

	if emailConfirmation.ExpiresAt.Before(time.Now().UTC()) {
		return nil, &errors.ErrConfirmationTokenAlreadyExpired
	}

	user.IsEmailConfirmed = true
	emailConfirmation.IsUsed = true

	if _, err := cm.ApplicationUserRepository.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	if err := cm.EmailConfirmationRepository.UpdateEmailConfirmation(ctx, emailConfirmation); err != nil {
		return nil, err
	}

	return &Response{
		AuthorizationCode:   authorizationCode.ID.String(),
		RedirectUri:         request.RedirectUri,
		State:               request.State,
		CodeChallenge:       request.CodeChallenge,
		CodeChallengeMethod: request.CodeChallengeMethod,
		ResponseType:        request.ResponseType,
	}, nil
}
