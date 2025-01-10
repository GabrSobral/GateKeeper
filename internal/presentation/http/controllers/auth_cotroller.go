package controllers

import (
	"net/http"

	utils "github.com/guard-service/internal/presentation/http"

	confirmuseremail "github.com/guard-service/internal/application/services/authentication/confirm-user-email"
	externalloginprovider "github.com/guard-service/internal/application/services/authentication/external-login-provider"
	forgotpassword "github.com/guard-service/internal/application/services/authentication/forgot-password"
	resendemailconfirmation "github.com/guard-service/internal/application/services/authentication/resend-email-confirmation"
	resetpassword "github.com/guard-service/internal/application/services/authentication/reset-password"
	signin "github.com/guard-service/internal/application/services/authentication/sign-in-credential"
	signup "github.com/guard-service/internal/application/services/authentication/sign-up-credential"
)

type AuthController struct {
	SignInCredentialService        *signin.SignInService
	SignUpCredentialService        *signup.SignUpService
	ConfirmUserEmailService        *confirmuseremail.ConfirmUserEmail
	ResendEmailConfirmationService *resendemailconfirmation.ResendEmailConfirmation
	ExternalLoginService           *externalloginprovider.ExternalLoginProvider
	ResetPasswordService           *resetpassword.ResetPasswordService
	ForgotPasswordService          *forgotpassword.ForgotPasswordService
}

// Sign In with credentials controller
func (ac *AuthController) SignInAuthController(writter http.ResponseWriter, request *http.Request) {
	var signInRequest signin.Request

	if err := utils.ParseBodyToSchema(&signInRequest, request); err != nil {
		panic(err)
	}

	response, err := ac.SignInCredentialService.Handler(request.Context(), signInRequest)

	if err != nil {
		panic(err)
	}

	utils.SendJson(writter, response, http.StatusOK)
}

// Sign Up with credentials controller
func (ac *AuthController) SignUpAuthController(writter http.ResponseWriter, request *http.Request) {
	var signUpRequest signup.Request

	if err := utils.ParseBodyToSchema(&signUpRequest, request); err != nil {
		panic(err)
	}

	if err := ac.SignUpCredentialService.Handler(request.Context(), signUpRequest); err != nil {
		panic(err)
	}

	utils.SendJson(writter, nil, http.StatusNoContent)
}

// Confirm email controller
func (ac *AuthController) ConfirmEmailAuthController(writter http.ResponseWriter, request *http.Request) {
	var confirmEmailRequest confirmuseremail.Request

	if err := utils.ParseBodyToSchema(&confirmEmailRequest, request); err != nil {
		panic(err)
	}

	response, err := ac.ConfirmUserEmailService.Handler(request.Context(), confirmEmailRequest)

	if err != nil {
		panic(err)
	}

	utils.SendJson(writter, response, http.StatusOK)
}

// Resend Email Confirmation controller
func (ac *AuthController) ResendEmailConfirmationAuthController(writter http.ResponseWriter, request *http.Request) {
	var resendEmailConfirmationRequest resendemailconfirmation.Request

	if err := utils.ParseBodyToSchema(&resendEmailConfirmationRequest, request); err != nil {
		panic(err)
	}

	if err := ac.ResendEmailConfirmationService.Handler(request.Context(), resendEmailConfirmationRequest); err != nil {
		panic(err)
	}

	utils.SendJson(writter, nil, http.StatusNoContent)
}

// Confirm email controller
func (ac *AuthController) ExternalLoginAuthController(writter http.ResponseWriter, request *http.Request) {
	var externalLoginRequest externalloginprovider.Request

	if err := utils.ParseBodyToSchema(&externalLoginRequest, request); err != nil {
		panic(err)
	}

	response, err := ac.ExternalLoginService.Handler(request.Context(), externalLoginRequest)

	if err != nil {
		panic(err)
	}

	utils.SendJson(writter, response, http.StatusOK)
}

// Forgot Password controller
func (ac *AuthController) ResetPasswordAuthController(writter http.ResponseWriter, request *http.Request) {
	var resetPasswordRequest resetpassword.Request

	if err := utils.ParseBodyToSchema(&resetPasswordRequest, request); err != nil {
		panic(err)
	}

	if err := ac.ResetPasswordService.Handler(request.Context(), resetPasswordRequest); err != nil {
		panic(err)
	}

	utils.SendJson(writter, nil, http.StatusNoContent)
}

// Forgot Password controller
func (ac *AuthController) ForgotPasswordAuthController(writter http.ResponseWriter, request *http.Request) {
	var forgotPasswordRequest forgotpassword.Request

	if err := utils.ParseBodyToSchema(&forgotPasswordRequest, request); err != nil {
		panic(err)
	}

	if err := ac.ForgotPasswordService.Handler(request.Context(), forgotPasswordRequest); err != nil {
		panic(err)
	}

	utils.SendJson(writter, nil, http.StatusNoContent)
}
