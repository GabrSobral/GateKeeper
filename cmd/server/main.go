package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/guard-service/cmd/server/docs"
	confirmuseremail "github.com/guard-service/internal/application/services/authentication/confirm-user-email"
	externalloginprovider "github.com/guard-service/internal/application/services/authentication/external-login-provider"
	forgotpassword "github.com/guard-service/internal/application/services/authentication/forgot-password"
	resendemailconfirmation "github.com/guard-service/internal/application/services/authentication/resend-email-confirmation"
	resetpassword "github.com/guard-service/internal/application/services/authentication/reset-password"
	signin "github.com/guard-service/internal/application/services/authentication/sign-in-credential"
	signup "github.com/guard-service/internal/application/services/authentication/sign-up-credential"
	getuserbyemail "github.com/guard-service/internal/application/services/user/get-user-by-email"
	getuserbyid "github.com/guard-service/internal/application/services/user/get-user-by-id"
	"github.com/guard-service/internal/domain/entities"
	"github.com/guard-service/internal/infra/database"
	repository_handlers "github.com/guard-service/internal/infra/database/repositories/handlers"
	inmemory_repositories "github.com/guard-service/internal/infra/database/repositories/in-memory"
	mailservice "github.com/guard-service/internal/infra/mail-service"
	"github.com/guard-service/internal/presentation/http/controllers"
	http_middlewares "github.com/guard-service/internal/presentation/http/middlewares"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Guard Service API
//	@version		1
//	@description	This is the Guard Service API documentation.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// @host		localhost:8080
// @BasePath	/v1
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	store, err := database.NewConnectionPool()

	if err != nil {
		panic(err)
	}

	userRepository := repository_handlers.UserRepository{Store: store}

	// inMemoryUserRepository := inmemory_repositories.InMemoryUserRepository{Users: make(map[string]*entities.User)}
	inMemoryUserProfileRepository := inmemory_repositories.InMemoryUserProfileRepository{Users: make(map[string]*entities.UserProfile)}
	inMemoryRefreshTokenRepository := inmemory_repositories.InMemoryRefreshTokenRepository{RefreshTokens: make(map[string]*entities.RefreshToken)}
	inMemoryEmailConfirmationRepository := inmemory_repositories.InMemoryEmailConfirmationRepository{Emails: make(map[string]*entities.EmailConfirmation)}
	inMemoryExternalLoginRepository := inmemory_repositories.InMemoryExternalLoginRepository{Logins: make(map[string]*entities.ExternalLogin)}
	inMemoryPasswordResetRepository := inmemory_repositories.InMemoryPasswordResetRepository{PasswordTokens: make(map[string]*entities.PasswordResetToken)}

	mailservice := mailservice.MailService{}

	authController := controllers.AuthController{
		SignInCredentialService: &signin.SignInService{
			UserRepository:         userRepository,
			UserProfileRepository:  inMemoryUserProfileRepository,
			RefreshTokenRepository: inMemoryRefreshTokenRepository,
		},
		SignUpCredentialService: &signup.SignUpService{
			UserRepository:              userRepository,
			UserProfileRepository:       inMemoryUserProfileRepository,
			RefreshTokenRepository:      inMemoryRefreshTokenRepository,
			EmailConfirmationRepository: inMemoryEmailConfirmationRepository,
			MailService:                 &mailservice,
		},
		ConfirmUserEmailService: &confirmuseremail.ConfirmUserEmail{
			UserRepository:              userRepository,
			EmailConfirmationRepository: inMemoryEmailConfirmationRepository,
			UserProfileRepository:       inMemoryUserProfileRepository,
			RefreshTokenRepository:      inMemoryRefreshTokenRepository,
		},
		ResendEmailConfirmationService: &resendemailconfirmation.ResendEmailConfirmation{
			UserRepository:              userRepository,
			UserProfileRepository:       inMemoryUserProfileRepository,
			EmailConfirmationRepository: inMemoryEmailConfirmationRepository,
			MailService:                 &mailservice,
		},
		ExternalLoginService: &externalloginprovider.ExternalLoginProvider{
			UserRepository:          userRepository,
			UserProfileRepository:   inMemoryUserProfileRepository,
			ExternalLoginRepository: inMemoryExternalLoginRepository,
		},
		ResetPasswordService: &resetpassword.ResetPasswordService{
			UserRepository:          userRepository,
			RefreshTokenRepository:  inMemoryRefreshTokenRepository,
			PasswordResetRepository: inMemoryPasswordResetRepository,
		},
		ForgotPasswordService: &forgotpassword.ForgotPasswordService{
			UserRepository:          userRepository,
			PasswordResetRepository: inMemoryPasswordResetRepository,
			UserProfileRepository:   inMemoryUserProfileRepository,
			MailService:             &mailservice,
		},
	}

	userController := controllers.UserController{
		GetUserByEmailService: &getuserbyemail.GetUserByEmail{
			UserRepository:        userRepository,
			UserProfileRepository: inMemoryUserProfileRepository,
		},
		GetUserByIDService: &getuserbyid.GetUserByID{
			UserRepository:        userRepository,
			UserProfileRepository: inMemoryUserProfileRepository,
		},
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Use(http_middlewares.ErrorHandler)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // 5 minutes
	}))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	// Health check
	r.Get("/health", func(writter http.ResponseWriter, request *http.Request) {
		writter.Write([]byte("Healthy"))
	})

	// Routes v1
	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-in", authController.SignInAuthController)
			r.Post("/sign-up", authController.SignUpAuthController)
			r.Post("/confirm-email", authController.ConfirmEmailAuthController)
			r.Post("/reset-password", authController.ResetPasswordAuthController)
			r.Post("/forgot-password", authController.ForgotPasswordAuthController)
			r.Post("/confirm-email/resend", authController.ResendEmailConfirmationAuthController)

			r.Post("/external-provider", authController.ExternalLoginAuthController)
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(http_middlewares.JwtHandler)

			r.Get("/by-email/{email}", userController.GetUserByEmailController)
			r.Get("/by-id/{userID}", userController.GetUserByIDController)
		})
	})

	slog.Info("✅ Server is running on port 8080")

	http.ListenAndServe(":8080", r)
}
