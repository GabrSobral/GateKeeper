package errors

import (
	"net/http"
)

type CustomError struct {
	Name    string
	Code    int
	Message string
	Title   string
}

func (e *CustomError) Error() string {
	return e.Name
}

var (
	ErrUserNotFound           = CustomError{Name: "ErrUserNotFound", Code: http.StatusNotFound, Message: "User was not found in the system", Title: "User not found"}
	ErrEmailOrPasswordInvalid = CustomError{Name: "ErrEmailOrPasswordInvalid", Code: http.StatusBadRequest, Message: "E-mail/password is incorrect or invalid", Title: "Invalid e-mail or password"}
	ErrInvalidEmail           = CustomError{Name: "ErrInvalidEmail", Code: http.StatusBadRequest, Message: "Invalid e-mail address, please provide a valid e-mail address", Title: "Invalid e-mail"}
	ErrEmailNotConfirmed      = CustomError{Name: "ErrEmailNotConfirmed", Code: http.StatusBadRequest, Message: "E-mail not confirmed, please confirm your e-mail address to continue", Title: "E-mail not confirmed"}
	ErrUserNotActive          = CustomError{Name: "ErrUserNotActive", Code: http.StatusBadRequest, Message: "User not active, please contact support", Title: "User not active"}
	ErrUserAlreadyExists      = CustomError{Name: "ErrUserAlreadyExists", Code: http.StatusBadRequest, Message: "An user is already registered with this e-mail, try another e-mail", Title: "User already exists"}
	ErrUserSignUpWithSocial   = CustomError{Name: "ErrUserSignUpWithSocial", Code: http.StatusBadRequest, Message: "User signed up with social login, please use social login", Title: "User signed up with social login"}

	ErrEmailConfirmationIsInCoolDown   = CustomError{Name: "ErrEmailConfirmationIsInCoolDown", Code: http.StatusBadRequest, Message: "E-mail confirmation is in cool down yet, wait a few minutes and try again", Title: "E-mail confirmation is in cool down"}
	ErrEmailConfirmationNotFound       = CustomError{Name: "ErrEmailConfirmationNotFound", Code: http.StatusBadRequest, Message: "E-mail confirmation not found", Title: "E_mail confirmation not found"}
	ErrConfirmationTokenAlreadyExpired = CustomError{Name: "ErrConfirmationTokenAlreadyExpired", Code: http.StatusBadRequest, Message: "Confirmation token already expired, try generating another one", Title: "Confirmation token already expired"}
	ErrConfirmationTokenAlreadyUsed    = CustomError{Name: "ErrConfirmationTokenAlreadyUsed", Code: http.StatusBadRequest, Message: "Confirmation token already used", Title: "Confirmation token already used"}
	ErrConfirmationTokenInvalid        = CustomError{Name: "ErrConfirmationTokenInvalid", Code: http.StatusBadRequest, Message: "Confirmation token invalid", Title: "Confirmation token invalid"}

	ErrPasswordResetNotFound      = CustomError{Name: "ErrPasswordResetNotFound", Code: http.StatusNotFound, Message: "Password reset token not found", Title: "Password reset token not found"}
	ErrPasswordResetTokenMismatch = CustomError{Name: "ErrPasswordResetTokenMismatch", Code: http.StatusBadRequest, Message: "Password reset token mismatch", Title: "Password reset token mismatch"}
	ErrPasswordResetTokenExpired  = CustomError{Name: "ErrPasswordResetTokenExpired", Code: http.StatusBadRequest, Message: "Password reset token expired", Title: "Password reset token expired"}

	ErrInvalidHash         = CustomError{Name: "ErrInvalidHash", Code: http.StatusBadRequest, Message: "The encoded hash is invalid", Title: "Invalid hash"}
	ErrIncompatibleVersion = CustomError{Name: "ErrIncompatibleVersion", Code: http.StatusBadRequest, Message: "Incompatible version of the hash algorithm", Title: "Incompatible version"}

	ErrApplicationNotFound      = CustomError{Name: "ErrApplicationNotFound", Code: http.StatusNotFound, Message: "Application not found", Title: "Application not found"}
	ErrAplicationSecretNotFound = CustomError{Name: "ErrAplicationSecretNotFound", Code: http.StatusNotFound, Message: "Application secret not found", Title: "Application secret not found"}
	ErrInvalidClientSecret      = CustomError{Name: "ErrInvalidClientSecret", Code: http.StatusBadRequest, Message: "Invalid client secret", Title: "Invalid client secret"}
	ErrClientSecretExpired      = CustomError{Name: "ErrClientSecretExpired", Code: http.StatusBadRequest, Message: "Client secret expired", Title: "Client secret expired"}

	ErrOrganizationNotFound = CustomError{Name: "ErrOrganizationNotFound", Code: http.StatusNotFound, Message: "Organization not found", Title: "Organization not found"}

	ErrUserRoleNotFound = CustomError{Name: "ErrUserRoleNotFound", Code: http.StatusNotFound, Message: "User role not found", Title: "User role not found"}

	ErrAuthorizationCodeNotFound           = CustomError{Name: "ErrAuthorizationCodeNotFound", Code: http.StatusNotFound, Message: "Authorization code not found", Title: "Authorization code not found"}
	ErrAuthorizationCodeExpired            = CustomError{Name: "ErrAuthorizationCodeExpired", Code: http.StatusBadRequest, Message: "Authorization code expired", Title: "Authorization code expired"}
	ErrAuthorizationCodeInvalidRedirectURI = CustomError{Name: "ErrAuthorizationCodeInvalidRedirectURI", Code: http.StatusBadRequest, Message: "Invalid redirect URI", Title: "Invalid redirect URI"}
	ErrAuthorizationCodeInvalidClientID    = CustomError{Name: "ErrAuthorizationCodeInvalidClientID", Code: http.StatusBadRequest, Message: "Invalid client ID", Title: "Invalid client ID"}
	ErrAuthorizationCodeInvalidPKCE        = CustomError{Name: "ErrAuthorizationCodeInvalidPKCE", Code: http.StatusBadRequest, Message: "Invalid PKCE", Title: "Invalid PKCE"}

	ErrSessionCodeNotFound    = CustomError{Name: "ErrSessionCodeNotFound", Code: http.StatusNotFound, Message: "Session code not found", Title: "Session code not found"}
	ErrSessionCodeExpired     = CustomError{Name: "ErrSessionCodeExpired", Code: http.StatusBadRequest, Message: "Session code expired", Title: "Session code expired"}
	ErrSessionCodeAlreadyUsed = CustomError{Name: "ErrSessionCodeAlreadyUsed", Code: http.StatusBadRequest, Message: "Session code already used", Title: "Session code already used"}

	ErrEmailMfaCodeExpired  = CustomError{Name: "ErrEmailMfaCodeExpired", Code: http.StatusBadRequest, Message: "E-mail MFA code expired", Title: "E-mail MFA code expired"}
	ErrEmailMfaCodeNotFound = CustomError{Name: "ErrEmailMfaCodeNotFound", Code: http.StatusNotFound, Message: "E-mail MFA code invalid", Title: "E-mail MFA code not found"}
	ErrMfaEmailNotEnabled   = CustomError{Name: "ErrMfaEmailNotEnabled", Code: http.StatusBadRequest, Message: "MFA e-mail not enabled", Title: "MFA e-mail not enabled to user"}
)

var ErrorsList = map[string]CustomError{
	"ErrUserNotFound":                        ErrUserNotFound,
	"ErrEmailOrPasswordInvalid":              ErrEmailOrPasswordInvalid,
	"ErrInvalidEmail":                        ErrInvalidEmail,
	"ErrEmailNotConfirmed":                   ErrEmailNotConfirmed,
	"ErrUserNotActive":                       ErrUserNotActive,
	"ErrUserAlreadyExists":                   ErrUserAlreadyExists,
	"ErrEmailConfirmationIsInCoolDown":       ErrEmailConfirmationIsInCoolDown,
	"ErrEmailConfirmationNotFound":           ErrEmailConfirmationNotFound,
	"ErrConfirmationTokenAlreadyExpired":     ErrConfirmationTokenAlreadyExpired,
	"ErrConfirmationTokenAlreadyUsed":        ErrConfirmationTokenAlreadyUsed,
	"ErrConfirmationTokenInvalid":            ErrConfirmationTokenInvalid,
	"ErrPasswordResetNotFound":               ErrPasswordResetNotFound,
	"ErrPasswordResetTokenMismatch":          ErrPasswordResetTokenMismatch,
	"ErrPasswordResetTokenExpired":           ErrPasswordResetTokenExpired,
	"ErrInvalidHash":                         ErrInvalidHash,
	"ErrIncompatibleVersion":                 ErrIncompatibleVersion,
	"ErrUserSignUpWithSocial":                ErrUserSignUpWithSocial,
	"ErrApplicationNotFound":                 ErrApplicationNotFound,
	"ErrAplicationSecretNotFound":            ErrAplicationSecretNotFound,
	"ErrOrganizationNotFound":                ErrOrganizationNotFound,
	"ErrUserRoleNotFound":                    ErrUserRoleNotFound,
	"ErrAuthorizationCodeNotFound":           ErrAuthorizationCodeNotFound,
	"ErrAuthorizationCodeExpired":            ErrAuthorizationCodeExpired,
	"ErrInvalidClientSecret":                 ErrInvalidClientSecret,
	"ErrClientSecretExpired":                 ErrClientSecretExpired,
	"ErrAuthorizationCodeInvalidRedirectURI": ErrAuthorizationCodeInvalidRedirectURI,
	"ErrAuthorizationCodeInvalidClientID":    ErrAuthorizationCodeInvalidClientID,
	"ErrAuthorizationCodeInvalidPKCE":        ErrAuthorizationCodeInvalidPKCE,
	"ErrSessionCodeNotFound":                 ErrSessionCodeNotFound,
	"ErrSessionCodeExpired":                  ErrSessionCodeExpired,
	"ErrSessionCodeAlreadyUsed":              ErrSessionCodeAlreadyUsed,
	"ErrMfaEmailNotEnabled":                  ErrMfaEmailNotEnabled,
	"ErrEmailMfaCodeExpired":                 ErrEmailMfaCodeExpired,
	"ErrEmailMfaCodeNotFound":                ErrEmailMfaCodeNotFound,
}
