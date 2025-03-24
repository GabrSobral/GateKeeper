package http_controllers

import (
	"net/http"

	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gate-keeper/internal/application/services/authentication/authorize"
	changepassword "github.com/gate-keeper/internal/application/services/authentication/change-password"
	confirmuseremail "github.com/gate-keeper/internal/application/services/authentication/confirm-user-email"
	externalloginprovider "github.com/gate-keeper/internal/application/services/authentication/external-login-provider"
	forgotpassword "github.com/gate-keeper/internal/application/services/authentication/forgot-password"
	"github.com/gate-keeper/internal/application/services/authentication/login"
	resendemailconfirmation "github.com/gate-keeper/internal/application/services/authentication/resend-email-confirmation"
	resetpassword "github.com/gate-keeper/internal/application/services/authentication/reset-password"
	"github.com/gate-keeper/internal/application/services/authentication/session"
	signin "github.com/gate-keeper/internal/application/services/authentication/sign-in-credential"
	signup "github.com/gate-keeper/internal/application/services/authentication/sign-up-credential"
	verifymfa "github.com/gate-keeper/internal/application/services/authentication/verify-mfa"
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

// Sign In with credentials controller
func (ac *AuthController) AuthorizeController(writter http.ResponseWriter, request *http.Request) {
	var authorizeRequest authorize.Request

	if err := http_router.ParseBodyToSchema(&authorizeRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[authorize.Request, *authorize.Response, authorize.AuthorizeService]{
		DbPool:  ac.DbPool,
		New:     authorize.New,
		Request: authorizeRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

// Sign In with credentials controller
func (ac *AuthController) LoginController(writter http.ResponseWriter, request *http.Request) {
	var loginRequest login.Request

	if err := http_router.ParseBodyToSchema(&loginRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[login.Request, *login.Response, login.LoginService]{
		DbPool:  ac.DbPool,
		New:     login.New,
		Request: loginRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

func (ac *AuthController) VerifyMfaController(writter http.ResponseWriter, request *http.Request) {
	var verifyMfaRequest verifymfa.Request

	if err := http_router.ParseBodyToSchema(&verifyMfaRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[verifymfa.Request, *verifymfa.Response, verifymfa.VerifyMfaService]{
		DbPool:  ac.DbPool,
		New:     verifymfa.New,
		Request: verifyMfaRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

func (ac *AuthController) ChangePasswordController(writter http.ResponseWriter, request *http.Request) {
	var changePasswordRequest changepassword.Request

	if err := http_router.ParseBodyToSchema(&changePasswordRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[changepassword.Request, changepassword.ChangePasswordService]{
		DbPool:  ac.DbPool,
		New:     changepassword.New,
		Request: changePasswordRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}

func (ac *AuthController) GetSessionAuthController(writter http.ResponseWriter, request *http.Request) {
	authorizationHeader := request.Header.Get("Authorization")
	accessToken := authorizationHeader[len("Bearer "):]

	service := session.SessionService{}

	response, err := service.Handler(request.Context(), session.Request{
		AccessToken: accessToken,
	})

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

	params := repositories.ParamsRs[confirmuseremail.Request, *confirmuseremail.Response, confirmuseremail.ConfirmUserEmail]{
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
