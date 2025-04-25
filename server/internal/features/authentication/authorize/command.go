package authorize

import "github.com/google/uuid"

type Command struct {
	ApplicationID       uuid.UUID `json:"applicationId" validate:"required"`
	SessionCode         string    `json:"sessionCode" validate:"required"`
	Email               string    `json:"email" validate:"required,email"`
	CodeChallenge       string    `json:"codeChallenge" validate:"required"`
	CodeChallengeMethod string    `json:"codeChallengeMethod" validate:"required"`
	RedirectUri         string    `json:"redirectUri" validate:"required"`
	ResponseType        string    `json:"responseType" validate:"required"`
	Scope               string    `json:"scope"`
	State               string    `json:"state" validate:"required"`
}
