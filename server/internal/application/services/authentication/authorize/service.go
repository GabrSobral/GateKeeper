package authorize

import (
	"context"

	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
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
	AuthorizationCode   string `json:"authorizationCode"`
	RedirectUri         string `json:"redirectUri"`
	State               string `json:"state"`
	CodeChallenge       string `json:"codeChallenge"`
	CodeChallengeMethod string `json:"codeChallengeMethod"`
	ResponseType        string `json:"responseType"`
}

type AuthorizeService struct {
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	UserProfileRepository     repository_interfaces.IUserProfileRepository
	RefreshTokenRepository    repository_interfaces.IRefreshTokenRepository
	AuthozationCodeRepository repository_interfaces.IApplicationAuthorizationCodeRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &AuthorizeService{
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository:    repository_handlers.RefreshTokenRepository{Store: q},
		AuthozationCodeRepository: repository_handlers.ApplicationAuthorizationCodeRepository{Store: q},
	}
}

func (ss *AuthorizeService) Handler(ctx context.Context, request Request) (*Response, error) {
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

	if err := ss.AuthozationCodeRepository.RemoveAuthorizationCode(ctx, user.ID, request.ApplicationID); err != nil {
		return nil, err
	}

	if err := ss.AuthozationCodeRepository.AddAuthorizationCode(ctx, authorizationCode); err != nil {
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
