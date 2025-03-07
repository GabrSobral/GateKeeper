package entities

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationAuthorizationCode struct {
	ID                  uuid.UUID
	ApplicationID       uuid.UUID
	ExpiresAt           time.Time
	Code                string
	ApplicationUserId   uuid.UUID
	RedirectUri         string
	CodeChallenge       string
	CodeChallengeMethod string
}

func CreateApplicationAuthorizationCode(applicationID, applicationUserID uuid.UUID, redirectUri, codeChallenge, codeChallegeMethod string) (*ApplicationAuthorizationCode, error) {
	userId, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	return &ApplicationAuthorizationCode{
		ID:                  userId,
		ApplicationID:       applicationID,
		ExpiresAt:           time.Now().UTC().Add(time.Minute * 5), // 5 minutes
		Code:                GenerateRandomString(128),
		ApplicationUserId:   applicationUserID,
		RedirectUri:         redirectUri,
		CodeChallenge:       codeChallenge,
		CodeChallengeMethod: codeChallegeMethod,
	}, nil
}
