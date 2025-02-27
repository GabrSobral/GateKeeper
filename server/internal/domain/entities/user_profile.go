package entities

import "github.com/google/uuid"

type UserProfile struct {
	UserID      uuid.UUID
	DisplayName string
	FirstName   string
	LastName    string
	PhoneNumber *string
	Address     *string
	PhotoURL    *string
}

func NewUserProfile(userID uuid.UUID, firstName, lastName, displayName string, phoneNumber, address, photoURL *string) *UserProfile {
	return &UserProfile{
		UserID:      userID,
		DisplayName: displayName,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		PhotoURL:    photoURL,
	}
}
