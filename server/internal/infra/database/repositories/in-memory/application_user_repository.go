package inmemory_repositories

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	"github.com/google/uuid"
)

type InMemoryApplicationUserRepository struct {
	Users map[string]*entities.ApplicationUser
}

func (r InMemoryApplicationUserRepository) AddUser(ctx context.Context, newUser *entities.ApplicationUser) error {
	r.Users[newUser.ID.String()] = newUser

	return nil
}

type GetUserByEmailParams struct {
	Email string
}

func (r InMemoryApplicationUserRepository) DeleteApplicationUser(ctx context.Context, applicationID, userID uuid.UUID) error {
	delete(r.Users, userID.String())

	return nil
}

func (r InMemoryApplicationUserRepository) GetUsersByApplicationID(ctx context.Context, applicationID uuid.UUID, limit, offset int) (*repository_interfaces.ApplicationUsersData, error) {
	users := []repository_interfaces.ApplicationUsers{}

	for _, user := range r.Users {
		if user.ApplicationID == applicationID {
			// users = append(users, repository_interfaces.ApplicationUsers{
			// 	ID:          user.ID,
			// 	DisplayName: "",
			// 	Email:       user.Email,
			// 	Roles:       []repository_interfaces.ApplicationRolesData{},
			// })
		}
	}

	return &repository_interfaces.ApplicationUsersData{
		Data:       users,
		TotalCount: len(users),
	}, nil
}

func (r InMemoryApplicationUserRepository) GetUserByEmail(ctx context.Context, email string, applicationID uuid.UUID) (*entities.ApplicationUser, error) {
	for _, user := range r.Users {
		if user.Email == email && user.ApplicationID == applicationID {
			return user, nil
		}
	}

	return nil, nil
}

func (r InMemoryApplicationUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.ApplicationUser, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, nil
}

func (r InMemoryApplicationUserRepository) IsUserExistsByEmail(ctx context.Context, email string, applicationID uuid.UUID) (bool, error) {
	for _, user := range r.Users {
		if user.Email == email && user.ApplicationID == applicationID {
			return true, nil
		}
	}

	return false, nil
}

func (r InMemoryApplicationUserRepository) IsUserExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return true, nil
		}
	}

	return false, nil
}

func (r InMemoryApplicationUserRepository) UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error) {
	r.Users[user.ID.String()] = user

	return user, nil
}
