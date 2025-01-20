package entities

import "github.com/google/uuid"

type UserProfile struct {
	UserID      uuid.UUID
	FirstName   string
	LastName    string
	PhoneNumber *string
	Address     *string
	PhotoURL    *string
}

func NewUserProfile(userID uuid.UUID, firstName, lastName string, phoneNumber, address, photoURL *string) *UserProfile {
	return &UserProfile{
		UserID:      userID,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		PhotoURL:    photoURL,
	}
}
