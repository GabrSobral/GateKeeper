package signin

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
)

func handleAuthorizationCode(ctx context.Context, service *SignInService, request Request) (*entities.ApplicationAuthorizationCode, error) {
	authorizationCode, err := service.AuthozationCodeRepository.GetAuthorizationCodeById(ctx, request.AuthorizationCode)

	if err != nil {
		return nil, err
	}

	if authorizationCode == nil {
		return nil, &errors.ErrAuthorizationCodeNotFound
	}

	currentDate := time.Now().UTC()

	if authorizationCode.ExpiresAt.Before(currentDate) {
		return nil, &errors.ErrAuthorizationCodeExpired
	}

	if authorizationCode.RedirectUri != request.RedirectURI {
		return nil, &errors.ErrAuthorizationCodeInvalidRedirectURI
	}

	if authorizationCode.ApplicationID != request.ClientID {
		return nil, &errors.ErrAuthorizationCodeInvalidClientID
	}

	if !validatePKCE(request.CodeVerifier, authorizationCode.CodeChallenge) {
		return nil, &errors.ErrAuthorizationCodeInvalidPKCE
	}

	return authorizationCode, nil
}

func generateCodeChallenge(codeVerifier string) string {
	hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash[:])
	return codeChallenge
}

// ValidatePKCE validates the given code challenge against the code verifier
func validatePKCE(codeVerifier, codeChallenge string) bool {
	generatedChallenge := generateCodeChallenge(codeVerifier)
	return generatedChallenge == codeChallenge
}
