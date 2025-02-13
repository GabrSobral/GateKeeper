package confirmuseremail

import (
	"context"
	"time"

	signin "github.com/gate-keeper/internal/application/services/authentication/sign-in-credential"
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
	Token         string    `json:"token" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	ApplicationID uuid.UUID `json:"application_id" validate:"required"`
}

type ConfirmUserEmail struct {
	ApplicationUserRepository   repository_interfaces.IApplicationUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository
	RefreshTokenRepository      repository_interfaces.IRefreshTokenRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *signin.Response] {
	return &ConfirmUserEmail{
		ApplicationUserRepository:   repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:       repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository:      repository_handlers.RefreshTokenRepository{Store: q},
		EmailConfirmationRepository: repository_handlers.EmailConfirmationRepository{Store: q},
	}
}

func (cm *ConfirmUserEmail) Handler(ctx context.Context, request Request) (*signin.Response, error) {
	user, err := cm.ApplicationUserRepository.GetUserByEmail(ctx, request.Email, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
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

	if emailConfirmation.ExpiresAt.Before(time.Now()) {
		return nil, &errors.ErrConfirmationTokenAlreadyExpired
	}

	user.IsEmailConfirmed = true
	emailConfirmation.IsUsed = true

	cm.ApplicationUserRepository.UpdateUser(ctx, user)
	cm.EmailConfirmationRepository.UpdateEmailConfirmation(ctx, emailConfirmation)

	userProfile, err := cm.UserProfileRepository.GetUserById(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	jwtToken, err := application_utils.CreateToken(application_utils.JWTClaims{
		UserID:    user.ID,
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		Email:     user.Email,
	})

	if err != nil {
		return nil, err
	}

	currentDate := time.Now().UTC()
	futureDate := currentDate.Add(time.Hour * 24 * 7).UTC() // 7 days from now

	cm.RefreshTokenRepository.RevokeRefreshTokenFromUser(ctx, user.ID)
	refreshToken, err := entities.CreateRefreshToken(user.ID, 5, futureDate)

	if err != nil {
		return nil, err
	}

	return &signin.Response{
		User: signin.UserResponse{
			ID:        user.ID,
			FirstName: userProfile.FirstName,
			LastName:  userProfile.LastName,
			Email:     user.Email,
			PhotoURL:  userProfile.PhotoURL,
			CreatedAt: user.CreatedAt,
		},
		AccessToken:  jwtToken,
		RefreshToken: refreshToken.ID,
	}, nil
}
