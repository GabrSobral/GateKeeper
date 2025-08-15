package externalloginprovider

import (
	"context"
	"strings"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Handler struct {
	repository IRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Command, *Response] {
	return &Handler{
		repository: Repository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, command Command) (*Response, error) {
	externalLogin, err := s.repository.GetExternalLoginByProviderKey(ctx, command.Provider, command.ProviderKey)

	if err != nil {
		return nil, err
	}

	if externalLogin != nil {
		user, err := s.repository.GetUserByID(ctx, externalLogin.UserID)

		if err != nil {
			return nil, err
		}

		if user == nil {
			return nil, &errors.ErrUserNotFound
		}

		userProfile, err := s.repository.GetUserProfileByID(ctx, user.ID)

		if err != nil {
			return nil, err
		}

		return &Response{
			UserID:           user.ID,
			UserEmail:        strings.ToLower(command.Email),
			FirstName:        userProfile.FirstName,
			LastName:         userProfile.LastName,
			CreatedAt:        user.CreatedAt,
			IsEmailConfirmed: user.IsEmailConfirmed,
			UpdatedAt:        nil,
		}, nil
	}

	user, err := s.repository.GetUserByEmail(ctx, command.Email, command.ApplicationID)

	if err != nil {
		return nil, err
	}

	userProfile := &entities.UserProfile{}

	if user == nil {
		userID, err := uuid.NewV7()

		if err != nil {
			return nil, err
		}

		user = &entities.ApplicationUser{
			ID:               userID,
			Email:            command.Email,
			PasswordHash:     nil,
			CreatedAt:        time.Now(),
			UpdatedAt:        nil,
			IsActive:         true,
			IsEmailConfirmed: true,
			ApplicationID:    command.ApplicationID,
			ShouldChangePass: false,
		}

		if err = s.repository.AddUser(ctx, user); err != nil {
			return nil, err
		}

		userProfile = &entities.UserProfile{
			UserID:      user.ID,
			FirstName:   command.FirstName,
			LastName:    command.LastName,
			PhoneNumber: nil,
			Address:     nil,
			PhotoURL:    nil,
		}

		if err := s.repository.AddUserProfile(ctx, userProfile); err != nil {
			return nil, err
		}
	} else {
		userProfileData, err := s.repository.GetUserProfileByID(ctx, user.ID)

		if err != nil {
			return nil, err
		}

		userProfile = userProfileData
	}

	newExternalLogin := &entities.ExternalLogin{
		UserID:      user.ID,
		Email:       strings.ToLower(command.Email),
		Provider:    command.Provider,
		ProviderKey: command.ProviderKey,
	}

	if err = s.repository.AddExternalLogin(ctx, newExternalLogin); err != nil {
		return nil, err
	}

	return &Response{
		UserID:           user.ID,
		UserEmail:        user.Email,
		FirstName:        userProfile.FirstName,
		LastName:         userProfile.LastName,
		CreatedAt:        user.CreatedAt,
		IsEmailConfirmed: user.IsEmailConfirmed,
		UpdatedAt:        nil,
	}, nil
}
