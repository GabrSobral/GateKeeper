package http_controllers

import (
	"net/http"

	createrole "github.com/gate-keeper/internal/features/application-role/create-role"
	deleterole "github.com/gate-keeper/internal/features/application-role/delete-role"
	listroles "github.com/gate-keeper/internal/features/application-role/list-roles"
	createsecret "github.com/gate-keeper/internal/features/application-secret/create-secret"
	deletesecret "github.com/gate-keeper/internal/features/application-secret/delete-secret"
	createapplicationuser "github.com/gate-keeper/internal/features/application-user/create-application-user"
	deleteapplicationuser "github.com/gate-keeper/internal/features/application-user/delete-application-user"
	editapplicationuser "github.com/gate-keeper/internal/features/application-user/edit-application-user"
	getapplicationuserbyid "github.com/gate-keeper/internal/features/application-user/get-application-user-by-id"
	createapplication "github.com/gate-keeper/internal/features/application/create-application"
	getapplicationauthdata "github.com/gate-keeper/internal/features/application/get-application-auth-data"
	getapplicationbyid "github.com/gate-keeper/internal/features/application/get-application-by-id"
	listapplications "github.com/gate-keeper/internal/features/application/list-applications"
	removeapplication "github.com/gate-keeper/internal/features/application/remove-application"
	updateapplication "github.com/gate-keeper/internal/features/application/update-application"
	"github.com/gate-keeper/internal/features/authentication/authorize"
	changepassword "github.com/gate-keeper/internal/features/authentication/change-password"
	confirmuseremail "github.com/gate-keeper/internal/features/authentication/confirm-user-email"
	externalloginprovider "github.com/gate-keeper/internal/features/authentication/external-login-provider"
	forgotpassword "github.com/gate-keeper/internal/features/authentication/forgot-password"
	generateauthappsecret "github.com/gate-keeper/internal/features/authentication/generate-auth-app-secret"
	login "github.com/gate-keeper/internal/features/authentication/generate-auth-app-secret"
	resendemailconfirmation "github.com/gate-keeper/internal/features/authentication/generate-auth-app-secret"
	resetpassword "github.com/gate-keeper/internal/features/authentication/reset-password"
	"github.com/gate-keeper/internal/features/authentication/session"
	signincredential "github.com/gate-keeper/internal/features/authentication/sign-in-credential"
	signupcredential "github.com/gate-keeper/internal/features/authentication/sign-up-credential"
	verifymfa "github.com/gate-keeper/internal/features/authentication/verify-mfa"
	createorganization "github.com/gate-keeper/internal/features/organization/create-organization"
	listorganizations "github.com/gate-keeper/internal/features/organization/list-organizations"
	removeorganization "github.com/gate-keeper/internal/features/organization/remove-organization"
	http_middlewares "github.com/gate-keeper/internal/presentation/http/middlewares"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetHttpRoutes(pool *pgxpool.Pool) http.Handler {
	// authController := AuthController{DbPool: pool}
	// applicationUserController := ApplicationUserController{DbPool: pool}
	// applicationController := ApplicationController{DbPool: pool}
	// organizationController := OrganizationController{DbPool: pool}
	// applicationSecretController := ApplicationSecretController{DbPool: pool}

	listApplicationsEndpoint := listapplications.Endpoint{DbPool: pool}
	updateApplicationEndpoint := updateapplication.Endpoint{DbPool: pool}
	removeApplicationEndpoint := removeapplication.Endpoint{DbPool: pool}
	createApplicationEndpoint := createapplication.Endpoint{DbPool: pool}
	getApplicationByIdEndpoint := getapplicationbyid.Endpoint{DbPool: pool}
	getApplicationAuthDataEndpoint := getapplicationauthdata.Endpoint{DbPool: pool}

	listRolesEndpoint := listroles.Endpoint{DbPool: pool}
	createRoleEndpoint := createrole.Endpoint{DbPool: pool}
	deleteRoleEndpoint := deleterole.Endpoint{DbPool: pool}

	createEndpoint := createsecret.Endpoint{DbPool: pool}
	deleteSecretEndpoint := deletesecret.Endpoint{DbPool: pool}

	createOrganizationEndpoint := createorganization.Endpoint{DbPool: pool}
	listOrganizationsEndpoint := listorganizations.Endpoint{DbPool: pool}
	removeOrganizationEndpoint := removeorganization.Endpoint{DbPool: pool}

	createApplicationUserEndpoint := createapplicationuser.Endpoint{DbPool: pool}
	updateApplicationUserEndpoint := editapplicationuser.Endpoint{DbPool: pool}
	deleteApplicationUserEndpoint := deleteapplicationuser.Endpoint{DbPool: pool}
	getApplicationUserByIdEndpoint := getapplicationuserbyid.Endpoint{DbPool: pool}

	authorizeEndpoint := authorize.Endpoint{DbPool: pool}
	changePasswordEndpoint := changepassword.Endpoint{DbPool: pool}
	confirmUserEmailEndpoint := confirmuseremail.Endpoint{DbPool: pool}
	externalLoginProviderEndpoint := externalloginprovider.Endpoint{DbPool: pool}
	sessionEndpoint := session.Endpoint{DbPool: pool}
	forgotPasswordEndpoint := forgotpassword.Endpoint{DbPool: pool}
	generateAuthAppSecretEndpoint := generateauthappsecret.Endpoint{DbPool: pool}
	loginEndpoint := login.Endpoint{DbPool: pool}
	resendEmailConfirmationEndpoint := resendemailconfirmation.Endpoint{DbPool: pool}
	resetRepositoryEndpoint := resetpassword.Endpoint{DbPool: pool}
	signInCredentialEndpoint := signincredential.Endpoint{DbPool: pool}
	signUpCredentialEndpoint := signupcredential.Endpoint{DbPool: pool}
	verfifyMfaEndpoint := verifymfa.Endpoint{DbPool: pool}

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

				r.Get("/", sessionEndpoint.Http)
			})

			r.Post("/authorize", authorizeEndpoint.Http)
			r.Post("/sign-in", signInCredentialEndpoint.Http)
			r.Post("/login", loginEndpoint.Http)
			r.Post("/generate-auth-secret", generateAuthAppSecretEndpoint.Http)
			r.Post("/verify-mfa", verfifyMfaEndpoint.Http)
			r.Post("/sign-up", signUpCredentialEndpoint.Http)
			r.Post("/confirm-email", confirmUserEmailEndpoint.Http)
			r.Post("/reset-password", resetRepositoryEndpoint.Http)
			r.Post("/change-password", changePasswordEndpoint.Http)
			r.Post("/forgot-password", forgotPasswordEndpoint.Http)
			r.Post("/external-provider", externalLoginProviderEndpoint.Http)
			r.Post("/confirm-email/resend", resendEmailConfirmationEndpoint.Http)

			r.Get("/application/{applicationID}/auth-data", getApplicationAuthDataEndpoint.Http)
		})

		r.Route("/organizations", func(r chi.Router) {
			// r.Use(http_middlewares.JwtHandler)

			r.Get("/", listOrganizationsEndpoint.Http)
			r.Post("/", createOrganizationEndpoint.Http)

			r.Route("/{organizationID}", func(r chi.Router) {
				r.Delete("/", removeOrganizationEndpoint.Http)

				r.Route("/applications", func(r chi.Router) {
					r.Get("/", listApplicationsEndpoint.Http)
					r.Post("/", createApplicationEndpoint.Http)
					r.Put("/{applicationID}", updateApplicationEndpoint.Http)
					r.Get("/{applicationID}", getApplicationByIdEndpoint.Http)
					r.Delete("/{applicationID}", removeApplicationEndpoint.Http)

					r.Route("/{applicationID}/users", func(r chi.Router) {
						r.Post("/", createApplicationUserEndpoint.Http)
						r.Put("/{userID}", updateApplicationUserEndpoint.Http)
						r.Get("/{userID}", getApplicationUserByIdEndpoint.Http)
						r.Delete("/{userID}", deleteApplicationUserEndpoint.Http)
					})

					r.Route("/{applicationID}/roles", func(r chi.Router) {
						r.Get("/", listRolesEndpoint.Http)
						r.Post("/", createRoleEndpoint.Http)
						r.Delete("/{roleID}", deleteRoleEndpoint.Http)
					})

					r.Route("/{applicationID}/secrets", func(r chi.Router) {
						r.Post("/", createEndpoint.Http)
						r.Delete("/{secretID}", deleteSecretEndpoint.Http)
					})
				})
			})
		})
	})

	return r
}
