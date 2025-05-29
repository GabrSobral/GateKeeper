package login

import "github.com/google/uuid"

type Response struct {
	MfaID              *uuid.UUID `json:"mfaId"`
	MfaType            *int       `json:"mfaType"`
	SessionCode        *string    `json:"sessionCode"`
	ChangePasswordCode *string    `json:"changePasswordCode"`
	Message            string     `json:"message"`
	UserID             uuid.UUID  `json:"userId"`
}
