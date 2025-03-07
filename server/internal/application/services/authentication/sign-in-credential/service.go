package signin

import (
	"context"
	"log/slog"
	"time"

	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	AuthorizationCode uuid.UUID `json:"authorizationCode"`
	ClientSecret      string    `json:"clientSecret"`
	ClientID          uuid.UUID `json:"clientId"`
	CodeVerifier      string    `json:"codeVerifier"`
	RedirectURI       string    `json:"redirectUri"`
}

type Response struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken uuid.UUID    `json:"refreshToken"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"displayName"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhotoURL    *string   `json:"photoUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

type SignInService struct {
	ApplicationUserRepository   repository_interfaces.IApplicationUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	RefreshTokenRepository      repository_interfaces.IRefreshTokenRepository
	AuthozationCodeRepository   repository_interfaces.IApplicationAuthorizationCodeRepository
	ApplicationSecretRepository repository_interfaces.IApplicationSecretRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &SignInService{
		ApplicationUserRepository:   repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:       repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository:      repository_handlers.RefreshTokenRepository{Store: q},
		AuthozationCodeRepository:   repository_handlers.ApplicationAuthorizationCodeRepository{Store: q},
		ApplicationSecretRepository: repository_handlers.ApplicationSecretRepository{Store: q},
	}
}

func (ss *SignInService) Handler(ctx context.Context, request Request) (*Response, error) {
	authorizationCode, err := handleAuthorizationCode(ctx, ss, request)

	if err != nil {
		return nil, err
	}

	secrets, err := ss.ApplicationSecretRepository.ListSecretsFromApplication(ctx, authorizationCode.ApplicationID)

	if err != nil {
		return nil, err
	}

	isClientSecretValid, err := application_utils.VerifyClientSecret(request.ClientSecret, secrets)

	if err != nil {
		return nil, err
	}

	if !isClientSecretValid {
		return nil, &errors.ErrInvalidClientSecret
	}

	if err := ss.AuthozationCodeRepository.RemoveAuthorizationCode(ctx, authorizationCode.ApplicationUserId, authorizationCode.ApplicationID); err != nil {
		return nil, err
	}

	user, err := ss.ApplicationUserRepository.GetUserByID(ctx, authorizationCode.ApplicationUserId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	refreshToken, err := assignRefreshToken(ctx, ss, *user)

	userProfile, err := ss.UserProfileRepository.GetUserById(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	jwtToken, err := assignTokenParams(*userProfile, *user)

	if err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, "User signed in successfully")

	return &Response{
		User: UserResponse{
			ID:          user.ID,
			DisplayName: userProfile.DisplayName,
			FirstName:   userProfile.FirstName,
			LastName:    userProfile.LastName,
			Email:       user.Email,
			PhotoURL:    userProfile.PhotoURL,
			CreatedAt:   user.CreatedAt,
		},
		AccessToken:  jwtToken,
		RefreshToken: refreshToken.ID,
	}, nil
}
