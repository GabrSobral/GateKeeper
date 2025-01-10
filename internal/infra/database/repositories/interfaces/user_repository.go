package repository_interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
)

type IUserRepository interface {
	AddUser(ctx context.Context, newUser *entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	IsUserExistsByEmail(ctx context.Context, email string) (bool, error)
	IsUserExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}
