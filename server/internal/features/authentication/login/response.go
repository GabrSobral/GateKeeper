package login

import "github.com/google/uuid"

type Response struct {
	MfaEmailRequired   bool      `json:"mfaEmailRequired"`
	MfaAuthAppRequired bool      `json:"mfaAuthAppRequired"`
	SessionCode        *string   `json:"sessionCode"`
	ChangePasswordCode *string   `json:"changePasswordCode"`
	Message            string    `json:"message"`
	UserID             uuid.UUID `json:"userId"`
}
