package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IApplicationUserRepository interface {
	AddUser(ctx context.Context, newUser *entities.ApplicationUser) error
	GetUserByEmail(ctx context.Context, email string, applicationID uuid.UUID) (*entities.ApplicationUser, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.ApplicationUser, error)
	IsUserExistsByEmail(ctx context.Context, email string, applicationID uuid.UUID) (bool, error)
	IsUserExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error)
}
