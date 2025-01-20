package application_utils

import (
	"regexp"
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
