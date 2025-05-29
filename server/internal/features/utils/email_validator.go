package application_utils

import (
	"regexp"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
)

// EmailValidator validates the given email string using a regex pattern
func EmailValidator(email string) bool {
	// Define the regex pattern for validating email addresses
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	re := regexp.MustCompile(emailRegexPattern)

	// Validate the email string
	return re.MatchString(email)
}

func VerifyClientSecret(clientSecret string, secrets *[]entities.ApplicationSecret) (bool, error) {
	for _, secret := range *secrets {
		if secret.Value == clientSecret {

			if secret.ExpiresAt != nil {
				if secret.ExpiresAt.Before(secret.CreatedAt) {
					return false, &errors.ErrClientSecretExpired
				}
			}

			return true, nil
		}
	}

	return false, nil
}
