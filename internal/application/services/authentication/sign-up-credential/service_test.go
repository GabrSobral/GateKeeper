package signup_test

import (
	"context"
	"testing"

	signup "github.com/gate-keeper/internal/application/services/authentication/sign-up-credential"
	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	inmemory_repositories "github.com/gate-keeper/internal/infra/database/repositories/in-memory"
	mailservice "github.com/gate-keeper/internal/infra/mail-service"
)

func setupTest() (*inmemory_repositories.InMemoryUserRepository, *inmemory_repositories.InMemoryUserProfileRepository, *inmemory_repositories.InMemoryRefreshTokenRepository, *signup.SignUpService) {
	inMemoryUserRepository := inmemory_repositories.InMemoryUserRepository{Users: make(map[string]*entities.User)}
	inMemoryUserProfileRepository := inmemory_repositories.InMemoryUserProfileRepository{Users: make(map[string]*entities.UserProfile)}
	inMemoryRefreshTokenRepository := inmemory_repositories.InMemoryRefreshTokenRepository{RefreshTokens: make(map[string]*entities.RefreshToken)}
	inMemoryEmailConfirmationRepository := inmemory_repositories.InMemoryEmailConfirmationRepository{Emails: make(map[string]*entities.EmailConfirmation)}

	mailServiceMock := mailservice.MailServiceMock{}

	signUpService := signup.SignUpService{
		UserRepository:              inMemoryUserRepository,
		UserProfileRepository:       inMemoryUserProfileRepository,
		RefreshTokenRepository:      inMemoryRefreshTokenRepository,
		EmailConfirmationRepository: inMemoryEmailConfirmationRepository,
		MailService:                 &mailServiceMock,
	}

	return &inMemoryUserRepository, &inMemoryUserProfileRepository, &inMemoryRefreshTokenRepository, &signUpService
}

func TestSignUpCredentialService(t *testing.T) {
	t.Run("Should return an error when user already exists", func(t *testing.T) {
		ctx := context.Background()

		request := signup.Request{
			Email:     "test@test.com",
			Password:  "123",
			FirstName: "Test",
			LastName:  "Test",
		}

		_, _, _, signUpService := setupTest()

		signUpService.Handler(ctx, request)

		err := signUpService.Handler(ctx, request)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if err != &errors.ErrUserAlreadyExists {
			t.Errorf("Expected user already exists, got %v", err.Error())
		}
	})

	t.Run("Should run without any error", func(t *testing.T) {
		ctx := context.Background()

		inMemoryUserRepository, inMemoryUserProfileRepository, _, signUpService := setupTest()

		request := signup.Request{
			Email:     "test@email.com",
			Password:  "123",
			FirstName: "Test",
			LastName:  "Test",
		}

		err := signUpService.Handler(ctx, request)

		if err != nil {
			t.Errorf("Expected nil, got error %v", err)
		}

		if len(inMemoryUserProfileRepository.Users) != 1 {
			t.Errorf("Expected users profiles length == 1, got %v", len(inMemoryUserProfileRepository.Users))
		}

		if len(inMemoryUserRepository.Users) != 1 {
			t.Errorf("Expected users length == 1, got %v", len(inMemoryUserRepository.Users))
		}
	})

}
