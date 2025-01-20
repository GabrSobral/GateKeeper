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
	userController := UserController{DbPool: pool}
	applicationController := ApplicationController{DbPool: pool}
	groupController := GroupController{DbPool: pool}
	tenantController := TenantController{DbPool: pool}

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
			r.Post("/external-provider", authController.ExternalLoginAuthController)
			r.Post("/confirm-email/resend", authController.ResendEmailConfirmationAuthController)
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(http_middlewares.JwtHandler)

			r.Get("/by-email/{email}", userController.GetUserByEmailController)
			r.Get("/by-id/{userID}", userController.GetUserByIDController)
		})

		r.Route("/applications", func(r chi.Router) {
			r.Use(http_middlewares.JwtHandler)

			r.Post("/", applicationController.CreateApplication)
			r.Delete("/{applicationID}", applicationController.RemoveApplication)
		})

		r.Route("/groups", func(r chi.Router) {
			r.Use(http_middlewares.JwtHandler)

			r.Post("/", groupController.CreateGroup)
			r.Delete("/{groupID}", groupController.RemoveGroup)
		})

		r.Route("/tenants", func(r chi.Router) {
			r.Use(http_middlewares.JwtHandler)

			r.Post("/", tenantController.CreateTenant)
			r.Delete("/{tenantID}", tenantController.RemoveTenant)
		})
	})

	return r
}
