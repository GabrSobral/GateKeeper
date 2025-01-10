package getuserbyemail

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/errors"
	repository_interfaces "github.com/guard-service/internal/infra/database/repositories/interfaces"
)

type Request struct {
	Email string `json:"email"`
}

type Response struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	FirstName       string    `json:"first_name"`
	Lastname        string    `json:"last_name"`
	Address         *string   `json:"address"`
	PhotoURL        *string   `json:"photo_url"`
	IsEmailVerified bool      `json:"is_email_verified"`
}

type GetUserByEmail struct {
	UserRepository        repository_interfaces.IUserRepository
	UserProfileRepository repository_interfaces.IUserProfileRepository
}

func (s *GetUserByEmail) Handler(ctx context.Context, request Request) (*Response, error) {
	user, err := s.UserRepository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, nil
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	userProfile, err := s.UserProfileRepository.GetUserById(ctx, user.ID)

	if err != nil {
		return nil, nil
	}

	return &Response{
		ID:              user.ID,
		Email:           user.Email,
		FirstName:       userProfile.FirstName,
		Lastname:        userProfile.LastName,
		Address:         userProfile.Address,
		PhotoURL:        userProfile.PhotoURL,
		IsEmailVerified: user.IsEmailConfirmed,
	}, nil
}
