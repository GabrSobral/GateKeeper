package getuserbyid

import (
	"context"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	UserID uuid.UUID `json:"userId"` // This is the user ID
}

type Response struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	FirstName       string    `json:"firstName"`
	Lastname        string    `json:"lastName"`
	Address         *string   `json:"address"`
	PhotoURL        *string   `json:"photoUrl"`
	IsEmailVerified bool      `json:"isEmailVerified"`
}

type GetUserByID struct {
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	UserProfileRepository     repository_interfaces.IUserProfileRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &GetUserByID{
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
	}
}

func (s *GetUserByID) Handler(ctx context.Context, request Request) (*Response, error) {
	user, err := s.ApplicationUserRepository.GetUserByID(ctx, request.UserID)

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
