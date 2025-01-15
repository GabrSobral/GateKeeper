package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IUserProfileRepository interface {
	AddUserProfile(ctx context.Context, newUser *entities.UserProfile) error
	GetUserById(ctx context.Context, userID uuid.UUID) (*entities.UserProfile, error)
}
