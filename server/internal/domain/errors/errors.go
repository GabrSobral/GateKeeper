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
	ErrUserNotFound           = CustomError{Name: "ErrUserNotFound", Code: http.StatusNotFound, Message: "user was not found in the system", Title: "User not found"}
	ErrEmailOrPasswordInvalid = CustomError{Name: "ErrEmailOrPasswordInvalid", Code: http.StatusBadRequest, Message: "email/password is incorrect or invalid", Title: "Invalid email or password"}
	ErrInvalidEmail           = CustomError{Name: "ErrInvalidEmail", Code: http.StatusBadRequest, Message: "invalid email address, please provide a valid email address", Title: "Invalid email"}
	ErrEmailNotConfirmed      = CustomError{Name: "ErrEmailNotConfirmed", Code: http.StatusBadRequest, Message: "email not confirmed, please confirm your email address to continue", Title: "Email not confirmed"}
	ErrUserNotActive          = CustomError{Name: "ErrUserNotActive", Code: http.StatusBadRequest, Message: "user not active, please contact support", Title: "User not active"}
	ErrUserAlreadyExists      = CustomError{Name: "ErrUserAlreadyExists", Code: http.StatusBadRequest, Message: "a user is already registered with this e-mail, try another e-mail", Title: "User already exists"}
	ErrUserSignUpWithSocial   = CustomError{Name: "ErrUserSignUpWithSocial", Code: http.StatusBadRequest, Message: "user signed up with social login, please use social login", Title: "User signed up with social login"}

	ErrEmailConfirmationIsInCoolDown   = CustomError{Name: "ErrEmailConfirmationIsInCoolDown", Code: http.StatusBadRequest, Message: "email confirmation is in cool down yet, wait a few minutes and try again", Title: "Email confirmation is in cool down"}
	ErrEmailConfirmationNotFound       = CustomError{Name: "ErrEmailConfirmationNotFound", Code: http.StatusBadRequest, Message: "email confirmation not found", Title: "Email confirmation not found"}
	ErrConfirmationTokenAlreadyExpired = CustomError{Name: "ErrConfirmationTokenAlreadyExpired", Code: http.StatusBadRequest, Message: "confirmation token already expired", Title: "Confirmation token already expired"}
	ErrConfirmationTokenAlreadyUsed    = CustomError{Name: "ErrConfirmationTokenAlreadyUsed", Code: http.StatusBadRequest, Message: "confirmation token already used", Title: "Confirmation token already used"}
	ErrConfirmationTokenInvalid        = CustomError{Name: "ErrConfirmationTokenInvalid", Code: http.StatusBadRequest, Message: "confirmation token invalid", Title: "Confirmation token invalid"}

	ErrPasswordResetNotFound      = CustomError{Name: "ErrPasswordResetNotFound", Code: http.StatusNotFound, Message: "password reset token not found", Title: "Password reset token not found"}
	ErrPasswordResetTokenMismatch = CustomError{Name: "ErrPasswordResetTokenMismatch", Code: http.StatusBadRequest, Message: "password reset token mismatch", Title: "Password reset token mismatch"}
	ErrPasswordResetTokenExpired  = CustomError{Name: "ErrPasswordResetTokenExpired", Code: http.StatusBadRequest, Message: "password reset token expired", Title: "Password reset token expired"}

	ErrInvalidHash         = CustomError{Name: "ErrInvalidHash", Code: http.StatusBadRequest, Message: "the encoded hash is invalid", Title: "Invalid hash"}
	ErrIncompatibleVersion = CustomError{Name: "ErrIncompatibleVersion", Code: http.StatusBadRequest, Message: "incompatible version of the hash algorithm", Title: "Incompatible version"}

	ErrAplicationNotFound       = CustomError{Name: "ErrAplicationNotFound", Code: http.StatusNotFound, Message: "application not found", Title: "Application not found"}
	ErrAplicationSecretNotFound = CustomError{Name: "ErrAplicationSecretNotFound", Code: http.StatusNotFound, Message: "application secret not found", Title: "Application secret not found"}

	ErrOrganizationNotFound = CustomError{Name: "ErrOrganizationNotFound", Code: http.StatusNotFound, Message: "organization not found", Title: "Organization not found"}

	ErrUserRoleNotFound = CustomError{Name: "ErrUserRoleNotFound", Code: http.StatusNotFound, Message: "user role not found", Title: "User role not found"}
)

var ErrorsList = map[string]CustomError{
	"ErrUserNotFound":                    ErrUserNotFound,
	"ErrEmailOrPasswordInvalid":          ErrEmailOrPasswordInvalid,
	"ErrInvalidEmail":                    ErrInvalidEmail,
	"ErrEmailNotConfirmed":               ErrEmailNotConfirmed,
	"ErrUserNotActive":                   ErrUserNotActive,
	"ErrUserAlreadyExists":               ErrUserAlreadyExists,
	"ErrEmailConfirmationIsInCoolDown":   ErrEmailConfirmationIsInCoolDown,
	"ErrEmailConfirmationNotFound":       ErrEmailConfirmationNotFound,
	"ErrConfirmationTokenAlreadyExpired": ErrConfirmationTokenAlreadyExpired,
	"ErrConfirmationTokenAlreadyUsed":    ErrConfirmationTokenAlreadyUsed,
	"ErrConfirmationTokenInvalid":        ErrConfirmationTokenInvalid,
	"ErrPasswordResetNotFound":           ErrPasswordResetNotFound,
	"ErrPasswordResetTokenMismatch":      ErrPasswordResetTokenMismatch,
	"ErrPasswordResetTokenExpired":       ErrPasswordResetTokenExpired,
	"ErrInvalidHash":                     ErrInvalidHash,
	"ErrIncompatibleVersion":             ErrIncompatibleVersion,
	"ErrUserSignUpWithSocial":            ErrUserSignUpWithSocial,
	"ErrAplicationNotFound":              ErrAplicationNotFound,
	"ErrAplicationSecretNotFound":        ErrAplicationSecretNotFound,
	"ErrOrganizationNotFound":            ErrOrganizationNotFound,
	"ErrUserRoleNotFound":                ErrUserRoleNotFound,
}
