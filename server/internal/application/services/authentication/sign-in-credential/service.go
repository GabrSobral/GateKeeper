package signin

import (
	"context"
	"log/slog"
	"time"

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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken uuid.UUID    `json:"refreshToken"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	PhotoURL  *string   `json:"photoUrl"`
	CreatedAt time.Time `json:"createdAt"`
}

type SignInService struct {
	UserRepository         repository_interfaces.IUserRepository
	UserProfileRepository  repository_interfaces.IUserProfileRepository
	RefreshTokenRepository repository_interfaces.IRefreshTokenRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &SignInService{
		UserRepository:         repository_handlers.UserRepository{Store: q},
		UserProfileRepository:  repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository: repository_handlers.RefreshTokenRepository{Store: q},
	}
}

func (ss *SignInService) Handler(ctx context.Context, request Request) (*Response, error) {
	slog.InfoContext(ctx, "Trying to sign in user with email: %s", request.Email, nil)

	user, err := ss.UserRepository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, &errors.ErrUserNotFound
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

	currentDate := time.Now().UTC()
	futureDate := currentDate.Add(time.Hour * 24 * 7).UTC() // 7 days from now

	ss.RefreshTokenRepository.RevokeRefreshTokenFromUser(ctx, user.ID)
	refreshToken, err := entities.CreateRefreshToken(user.ID, 5, futureDate)

	ss.RefreshTokenRepository.AddRefreshToken(ctx, refreshToken)

	if err != nil {
		return nil, err
	}

	userProfile, err := ss.UserProfileRepository.GetUserById(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	jwtToken, err := application_utils.CreateToken(
		application_utils.JWTClaims{
			UserID:    user.ID,
			Email:     user.Email,
			FirstName: userProfile.FirstName,
			LastName:  userProfile.LastName,
		},
	)

	if err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, "User signed in successfully")

	return &Response{
		User: UserResponse{
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
