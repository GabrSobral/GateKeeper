package http_controllers

import (
	"net/http"

	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/jackc/pgx/v5/pgxpool"

	confirmuseremail "github.com/gate-keeper/internal/application/services/authentication/confirm-user-email"
	externalloginprovider "github.com/gate-keeper/internal/application/services/authentication/external-login-provider"
	forgotpassword "github.com/gate-keeper/internal/application/services/authentication/forgot-password"
	resendemailconfirmation "github.com/gate-keeper/internal/application/services/authentication/resend-email-confirmation"
	resetpassword "github.com/gate-keeper/internal/application/services/authentication/reset-password"
	signin "github.com/gate-keeper/internal/application/services/authentication/sign-in-credential"
	signup "github.com/gate-keeper/internal/application/services/authentication/sign-up-credential"
)

type AuthController struct {
	DbPool *pgxpool.Pool
}

// Sign In with credentials controller
func (ac *AuthController) SignInAuthController(writter http.ResponseWriter, request *http.Request) {
	var signInRequest signin.Request

	if err := http_router.ParseBodyToSchema(&signInRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[signin.Request, *signin.Response, signin.SignInService]{
		DbPool:  ac.DbPool,
		New:     signin.New,
		Request: signInRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

// Sign Up with credentials controller
func (ac *AuthController) SignUpAuthController(writter http.ResponseWriter, request *http.Request) {
	var signUpRequest signup.Request

	if err := http_router.ParseBodyToSchema(&signUpRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[signup.Request, signup.SignUpService]{
		DbPool:  ac.DbPool,
		New:     signup.New,
		Request: signUpRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}

// Confirm email controller
func (ac *AuthController) ConfirmEmailAuthController(writter http.ResponseWriter, request *http.Request) {
	var confirmEmailRequest confirmuseremail.Request

	if err := http_router.ParseBodyToSchema(&confirmEmailRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[confirmuseremail.Request, *signin.Response, confirmuseremail.ConfirmUserEmail]{
		DbPool:  ac.DbPool,
		New:     confirmuseremail.New,
		Request: confirmEmailRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

// Resend Email Confirmation controller
func (ac *AuthController) ResendEmailConfirmationAuthController(writter http.ResponseWriter, request *http.Request) {
	var resendEmailConfirmationRequest resendemailconfirmation.Request

	if err := http_router.ParseBodyToSchema(&resendEmailConfirmationRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[resendemailconfirmation.Request, resendemailconfirmation.ResendEmailConfirmation]{
		DbPool:  ac.DbPool,
		New:     resendemailconfirmation.New,
		Request: resendEmailConfirmationRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}

// Confirm email controller
func (ac *AuthController) ExternalLoginAuthController(writter http.ResponseWriter, request *http.Request) {
	var externalLoginRequest externalloginprovider.Request

	if err := http_router.ParseBodyToSchema(&externalLoginRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[externalloginprovider.Request, *externalloginprovider.Response, externalloginprovider.ExternalLoginProvider]{
		DbPool:  ac.DbPool,
		New:     externalloginprovider.New,
		Request: externalLoginRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

// Forgot Password controller
func (ac *AuthController) ResetPasswordAuthController(writter http.ResponseWriter, request *http.Request) {
	var resetPasswordRequest resetpassword.Request

	if err := http_router.ParseBodyToSchema(&resetPasswordRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[resetpassword.Request, resetpassword.ResetPasswordService]{
		DbPool:  ac.DbPool,
		New:     resetpassword.New,
		Request: resetPasswordRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}

// Forgot Password controller
func (ac *AuthController) ForgotPasswordAuthController(writter http.ResponseWriter, request *http.Request) {
	var forgotPasswordRequest forgotpassword.Request

	if err := http_router.ParseBodyToSchema(&forgotPasswordRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[forgotpassword.Request, forgotpassword.ForgotPasswordService]{
		DbPool:  ac.DbPool,
		New:     forgotpassword.New,
		Request: forgotPasswordRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}
