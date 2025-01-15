package inmemory_repositories

import (
	"context"
	"errors"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type InMemoryUserProfileRepository struct {
	Users map[string]*entities.UserProfile
}

func (r InMemoryUserProfileRepository) AddUserProfile(ctx context.Context, newUserProfile *entities.UserProfile) error {
	r.Users[newUserProfile.UserID.String()] = newUserProfile

	return nil
}

func (r InMemoryUserProfileRepository) GetUserById(ctx context.Context, userID uuid.UUID) (*entities.UserProfile, error) {
	userProfile, ok := r.Users[userID.String()]

	if !ok {
		return nil, errors.New("user profile not found")
	}

	return userProfile, nil
}
