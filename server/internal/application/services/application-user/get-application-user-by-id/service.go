package getapplicationuserbyid

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
	ID                  uuid.UUID          `json:"id"`
	Email               string             `json:"email"`
	DisplayName         string             `json:"displayName"`
	IsActive            bool               `json:"isActive"`
	FirstName           string             `json:"firstName"`
	Lastname            string             `json:"lastName"`
	Address             *string            `json:"address"`
	PhotoURL            *string            `json:"photoUrl"`
	IsMfaEmailEnabled   bool               `json:"isMfaEmailEnabled"`
	IsMfaAuthAppEnabled bool               `json:"isMfaAuthAppEnabled"`
	IsEmailVerified     bool               `json:"isEmailVerified"`
	Badges              []UserRoleResponse `json:"badges"`
}

type UserRoleResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetApplicationUserByID struct {
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	UserProfileRepository     repository_interfaces.IUserProfileRepository
	UserRoleRepository        repository_interfaces.IUserRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &GetApplicationUserByID{
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
		UserRoleRepository:        repository_handlers.UserRoleRepository{Store: q},
	}
}

func (s *GetApplicationUserByID) Handler(ctx context.Context, request Request) (*Response, error) {
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

	badges, err := s.UserRoleRepository.GetRolesByUserID(ctx, user.ID)

	if err != nil {
		return nil, nil
	}

	badgesResponse := make([]UserRoleResponse, 0)

	for _, badge := range badges {
		badgesResponse = append(badgesResponse, UserRoleResponse{
			ID:   badge.ID,
			Name: badge.Name,
		})
	}

	return &Response{
		ID:                  user.ID,
		Email:               user.Email,
		IsActive:            user.IsActive,
		DisplayName:         userProfile.DisplayName,
		FirstName:           userProfile.FirstName,
		Lastname:            userProfile.LastName,
		Address:             userProfile.Address,
		PhotoURL:            userProfile.PhotoURL,
		IsMfaEmailEnabled:   user.IsMfaEmailEnabled,
		IsMfaAuthAppEnabled: user.IsMfaAuthAppEnabled,
		IsEmailVerified:     user.IsEmailConfirmed,
		Badges:              badgesResponse,
	}, nil
}
