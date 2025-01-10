package application_utils

import (
	"os"

	"golang.org/x/crypto/argon2"
)

// HashPassword hashes the given password using argon2
func HashPassword(password string) (string, error) {
	salt := os.Getenv("PASSWORD_SALT")
	hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)

	return string(hash), nil
}

// ComparePassword compares the hashed password with the normal password
func ComparePassword(hashedPassword *string, normalPassword string) (bool, error) {
	password, err := HashPassword(normalPassword)

	if err != nil {
		return false, err
	}

	return *hashedPassword == password, nil
}
