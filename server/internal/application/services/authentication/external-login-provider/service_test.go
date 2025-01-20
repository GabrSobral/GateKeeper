package externalloginprovider_test

import (
	"context"
	"testing"

	externalloginprovider "github.com/gate-keeper/internal/application/services/authentication/external-login-provider"
	"github.com/gate-keeper/internal/domain/entities"
	inmemory_repositories "github.com/gate-keeper/internal/infra/database/repositories/in-memory"
)

func TestExternalLoginService(t *testing.T) {
	t.Run("Should run without any error when is the first access", func(t *testing.T) {
		ctx := context.Background()

		inMemoryUserRepository := inmemory_repositories.InMemoryUserRepository{Users: make(map[string]*entities.User)}
		inMemoryUserProfileRepository := inmemory_repositories.InMemoryUserProfileRepository{Users: make(map[string]*entities.UserProfile)}
		inMemoryExternalLoginRepository := inmemory_repositories.InMemoryExternalLoginRepository{Logins: make(map[string]*entities.ExternalLogin)}

		externalLoginService := externalloginprovider.ExternalLoginProvider{
			UserRepository:          inMemoryUserRepository,
			UserProfileRepository:   inMemoryUserProfileRepository,
			ExternalLoginRepository: inMemoryExternalLoginRepository,
		}

		request := externalloginprovider.Request{
			Provider:    "google",
			ProviderKey: "1234567890",
			Email:       "test@email.com",
			FirstName:   "Test",
			LastName:    "User",
		}

		response, err := externalLoginService.Handler(ctx, request)

		if err != nil {
			t.Errorf("Error should be nil, got %v", err)
		}

		if response == nil {
			t.Error("Response should not be nil")
		}

		if len(inMemoryUserRepository.Users) == 0 {
			t.Error("User should be created")
		}

		if len(inMemoryUserProfileRepository.Users) == 0 {
			t.Error("User profile should be created")
		}

		if len(inMemoryExternalLoginRepository.Logins) == 0 {
			t.Error("External login should be created")
		}
	})

	t.Run("Should run without any error when is the second access or more", func(t *testing.T) {
		ctx := context.Background()

		inMemoryUserRepository := inmemory_repositories.InMemoryUserRepository{Users: make(map[string]*entities.User)}
		inMemoryUserProfileRepository := inmemory_repositories.InMemoryUserProfileRepository{Users: make(map[string]*entities.UserProfile)}
		inMemoryExternalLoginRepository := inmemory_repositories.InMemoryExternalLoginRepository{Logins: make(map[string]*entities.ExternalLogin)}

		externalLoginService := externalloginprovider.ExternalLoginProvider{
			UserRepository:          inMemoryUserRepository,
			UserProfileRepository:   inMemoryUserProfileRepository,
			ExternalLoginRepository: inMemoryExternalLoginRepository,
		}

		request := externalloginprovider.Request{
			Provider:    "google",
			ProviderKey: "1234567890",
			Email:       "test@email.com",
			FirstName:   "Test",
			LastName:    "User",
		}

		response, err := externalLoginService.Handler(ctx, request)

		if err != nil {
			t.Errorf("Error should be nil, got %v", err)
		}

		if response == nil {
			t.Error("Response should not be nil")
		}

		if len(inMemoryUserRepository.Users) == 0 {
			t.Error("User should be created")
		}

		if len(inMemoryUserProfileRepository.Users) == 0 {
			t.Error("User profile should be created")
		}

		if len(inMemoryExternalLoginRepository.Logins) == 0 {
			t.Error("External login should be created")
		}

		response, err = externalLoginService.Handler(ctx, request)

		if err != nil {
			t.Errorf("Error should be nil, got %v", err)
		}

		if response == nil {
			t.Error("Response should not be nil")
		}

		if len(inMemoryUserRepository.Users) != 1 {
			t.Error("Users Length should be 1")
		}

		if len(inMemoryUserProfileRepository.Users) != 1 {
			t.Error("Length should be 1")
		}

		if len(inMemoryExternalLoginRepository.Logins) != 1 {
			t.Error("Length should be 1")
		}
	})
}
