package inmemory_repositories

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type InMemoryUserRepository struct {
	Users map[string]*entities.User
}

func (r InMemoryUserRepository) AddUser(ctx context.Context, newUser *entities.User) error {
	r.Users[newUser.ID.String()] = newUser

	return nil
}

type GetUserByEmailParams struct {
	Email string
}

func (r InMemoryUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	for _, user := range r.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, nil
}

func (r InMemoryUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, nil
}

func (r InMemoryUserRepository) IsUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	for _, user := range r.Users {
		if user.Email == email {
			return true, nil
		}
	}

	return false, nil
}

func (r InMemoryUserRepository) IsUserExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return true, nil
		}
	}

	return false, nil
}

func (r InMemoryUserRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	r.Users[user.ID.String()] = user

	return user, nil
}
