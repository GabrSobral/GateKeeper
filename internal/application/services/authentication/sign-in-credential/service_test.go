package signin_test

import (
	"context"
	"testing"
	"time"

	signin "github.com/gate-keeper/internal/application/services/authentication/sign-in-credential"
	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	inmemory_repositories "github.com/gate-keeper/internal/infra/database/repositories/in-memory"
	"github.com/google/uuid"
)

func setupTest() (*inmemory_repositories.InMemoryUserRepository, *inmemory_repositories.InMemoryUserProfileRepository, *inmemory_repositories.InMemoryRefreshTokenRepository, *signin.SignInService) {
	inMemoryUserRepository := inmemory_repositories.InMemoryUserRepository{Users: make(map[string]*entities.User)}
	inMemoryUserProfileRepository := inmemory_repositories.InMemoryUserProfileRepository{Users: make(map[string]*entities.UserProfile)}
	inMemoryRefreshTokenRepository := inmemory_repositories.InMemoryRefreshTokenRepository{RefreshTokens: make(map[string]*entities.RefreshToken)}

	signInService := signin.SignInService{
		UserRepository:         inMemoryUserRepository,
		UserProfileRepository:  inMemoryUserProfileRepository,
		RefreshTokenRepository: inMemoryRefreshTokenRepository,
	}

	return &inMemoryUserRepository, &inMemoryUserProfileRepository, &inMemoryRefreshTokenRepository, &signInService
}

func TestSignInCredentialService(t *testing.T) {
	t.Run("Should return an error when user is not found", func(t *testing.T) {
		ctx := context.Background()

		request := signin.Request{
			Email:    "test@test.com",
			Password: "123",
		}

		_, _, _, signInService := setupTest()

		response, err := signInService.Handler(ctx, request)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if response != nil {
			t.Errorf("Expected nil, got %v", response)
		}

		if err != &errors.ErrUserNotFound {
			t.Errorf("Expected user not found, got %v", err.Error())
		}
	})

	t.Run("Should return an error when email is not valid", func(t *testing.T) {
		ctx := context.Background()

		inMemoryUserRepository, inMemoryUserProfileRepository, _, signInService := setupTest()

		hashedPassword, err := application_utils.HashPassword("123")

		if err != nil {
			t.Errorf("Expected nil, got error %v", err)
		}

		userId, err := uuid.NewV7()

		if err != nil {
			t.Errorf("Expected nil, got error %v", err)
		}

		inMemoryUserRepository.Users[userId.String()] = &entities.User{
			ID:               userId,
			Email:            "test@email.com",
			PasswordHash:     &hashedPassword,
			IsActive:         true,
			CreatedAt:        time.Now().UTC(),
			IsEmailConfirmed: true,
		}

		inMemoryUserProfileRepository.Users[userId.String()] = &entities.UserProfile{
			UserID:      userId,
			FirstName:   "Test",
			LastName:    "Test",
			PhoneNumber: nil,
			Address:     nil,
			PhotoURL:    nil,
		}

		request := signin.Request{
			Email:    "test@email.com",
			Password: "123",
		}

		response, err := signInService.Handler(ctx, request)

		if err != nil {
			t.Errorf("Expected nil, got error %v", err)
		}

		if response == nil {
			t.Errorf("Expected response, got nil")
		}
	})

	t.Run("Should run without any error", func(t *testing.T) {
		ctx := context.Background()

		inMemoryUserRepository, inMemoryUserProfileRepository, inMemoryRefreshTokenRepository, signInService := setupTest()

		hashedPassword, err := application_utils.HashPassword("123")

		if err != nil {
			t.Errorf("Expected nil, got error %v", err)
		}

		newId, _ := uuid.NewV7()

		inMemoryUserRepository.Users[newId.String()] = &entities.User{
			ID:               newId,
			Email:            "test@email.com",
			PasswordHash:     &hashedPassword,
			IsActive:         true,
			CreatedAt:        time.Now().UTC(),
			IsEmailConfirmed: true,
		}

		inMemoryUserProfileRepository.Users[newId.String()] = &entities.UserProfile{
			UserID:      newId,
			FirstName:   "Test",
			LastName:    "Test",
			PhoneNumber: nil,
			Address:     nil,
			PhotoURL:    nil,
		}

		request := signin.Request{
			Email:    "test@email.com",
			Password: "123",
		}

		response, err := signInService.Handler(ctx, request)

		if err != nil {
			t.Errorf("Expected nil, got error %v", err)
		}

		if response == nil {
			t.Errorf("Expected response, got nil")
		}

		if len(inMemoryRefreshTokenRepository.RefreshTokens) == 0 {
			t.Errorf("Expected refresh token length > 0, got %v", len(inMemoryRefreshTokenRepository.RefreshTokens))
		}
	})

}
