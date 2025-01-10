package confirmuseremail

import (
	"context"
	"time"

	signin "github.com/guard-service/internal/application/services/authentication/sign-in-credential"
	application_utils "github.com/guard-service/internal/application/utils"
	"github.com/guard-service/internal/domain/entities"
	"github.com/guard-service/internal/domain/errors"
	repository_interfaces "github.com/guard-service/internal/infra/database/repositories/interfaces"
)

type Request struct {
	Token string `json:"token" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type ConfirmUserEmail struct {
	UserRepository              repository_interfaces.IUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository
	RefreshTokenRepository      repository_interfaces.IRefreshTokenRepository
}

func (cm *ConfirmUserEmail) Handler(ctx context.Context, request Request) (*signin.Response, error) {
	user, err := cm.UserRepository.GetUserByEmail(ctx, request.Email)

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

	cm.UserRepository.UpdateUser(ctx, user)
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
