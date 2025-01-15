package main

import (
	"log/slog"
	"net/http"

	_ "github.com/gate-keeper/cmd/server/docs"
	"github.com/gate-keeper/internal/infra/database"
	"github.com/gate-keeper/internal/presentation/http/controllers"
	http_middlewares "github.com/gate-keeper/internal/presentation/http/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			GateKeeper API
//	@version		1
//	@description	This is the GateKeeper API documentation.
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//
// @host		localhost:8080
// @BasePath	/v1
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	pool, err := database.NewConnectionPool()

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	authController := controllers.AuthController{DbPool: pool}
	userController := controllers.UserController{DbPool: pool}

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
