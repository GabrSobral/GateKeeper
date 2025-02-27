package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type ApplicationUsersData struct {
	TotalCount int                `json:"totalCount"`
	Data       []ApplicationUsers `json:"data"`
}

type ApplicationRolesData struct {
	TotalCount int                `json:"totalCount"`
	Data       []ApplicationRoles `json:"data"`
}

type ApplicationUsers struct {
	ID          uuid.UUID          `json:"id"`
	DisplayName string             `json:"displayName"`
	Email       string             `json:"email"`
	Roles       []ApplicationRoles `json:"roles"`
}

type ApplicationRoles struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

type IApplicationUserRepository interface {
	AddUser(ctx context.Context, newUser *entities.ApplicationUser) error
	GetUserByEmail(ctx context.Context, email string, applicationID uuid.UUID) (*entities.ApplicationUser, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.ApplicationUser, error)
	IsUserExistsByEmail(ctx context.Context, email string, applicationID uuid.UUID) (bool, error)
	IsUserExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error)
	GetUsersByApplicationID(ctx context.Context, applicationID uuid.UUID, limit, offset int) (*ApplicationUsersData, error)
	DeleteApplicationUser(ctx context.Context, applicationID, userID uuid.UUID) error
}
