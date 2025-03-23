package http_controllers

import (
	"net/http"

	http_middlewares "github.com/gate-keeper/internal/presentation/http/middlewares"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetHttpRoutes(pool *pgxpool.Pool) http.Handler {
	authController := AuthController{DbPool: pool}
	applicationUserController := ApplicationUserController{DbPool: pool}
	applicationController := ApplicationController{DbPool: pool}
	organizationController := OrganizationController{DbPool: pool}
	applicationRoleController := ApplicationRoleController{DbPool: pool}
	applicationSecretController := ApplicationSecretController{DbPool: pool}

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
			r.Route("/session", func(r chi.Router) {
				r.Use(http_middlewares.JwtHandler)

				r.Get("/", authController.GetSessionAuthController)
			})

			r.Post("/authorize", authController.AuthorizeController)
			r.Post("/sign-in", authController.SignInAuthController)
			r.Post("/login", authController.LoginController)
			r.Post("/verify-mfa", authController.VerifyMfaController)
			r.Post("/sign-up", authController.SignUpAuthController)
			r.Post("/confirm-email", authController.ConfirmEmailAuthController)
			r.Post("/reset-password", authController.ResetPasswordAuthController)
			r.Post("/forgot-password", authController.ForgotPasswordAuthController)
			r.Post("/external-provider", authController.ExternalLoginAuthController)
			r.Post("/confirm-email/resend", authController.ResendEmailConfirmationAuthController)
			r.Get("/application/{applicationID}/auth-data", applicationController.GetApplicationAuthData)
		})

		r.Route("/organizations", func(r chi.Router) {
			// r.Use(http_middlewares.JwtHandler)

			r.Get("/", organizationController.ListOrganizations)
			r.Post("/", organizationController.CreateOrganization)

			r.Route("/{organizationID}", func(r chi.Router) {
				r.Delete("/", organizationController.RemoveOrganization)

				r.Route("/applications", func(r chi.Router) {
					r.Get("/", applicationController.ListApplications)
					r.Post("/", applicationController.CreateApplication)
					r.Put("/{applicationID}", applicationController.UpdateApplication)
					r.Get("/{applicationID}", applicationController.GetApplicationByID)
					r.Delete("/{applicationID}", applicationController.RemoveApplication)

					r.Route("/{applicationID}/users", func(r chi.Router) {
						r.Post("/", applicationUserController.CreateUser)
						r.Put("/{userID}", applicationUserController.UpdateUser)
						r.Delete("/{userID}", applicationUserController.DeleteUser)
						r.Get("/{userID}", applicationUserController.GetUserByIDController)
					})

					r.Route("/{applicationID}/roles", func(r chi.Router) {
						r.Get("/", applicationRoleController.ListRoles)
						r.Post("/", applicationRoleController.CreateRole)
						r.Delete("/{roleID}", applicationRoleController.RemoveRole)
					})

					r.Route("/{applicationID}/secrets", func(r chi.Router) {
						r.Post("/", applicationSecretController.CreateSecret)
						r.Delete("/{secretID}", applicationSecretController.RemoveSecret)
					})
				})
			})
		})
	})

	return r
}
