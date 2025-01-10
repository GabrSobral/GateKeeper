package repository_interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
)

type IUserProfileRepository interface {
	AddUserProfile(ctx context.Context, newUser *entities.UserProfile) error
	GetUserById(ctx context.Context, userID uuid.UUID) (*entities.UserProfile, error)
}
