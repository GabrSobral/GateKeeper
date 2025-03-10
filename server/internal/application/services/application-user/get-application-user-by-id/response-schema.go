package getapplicationuserbyid

import "github.com/google/uuid"

type Response struct {
	ID                  uuid.UUID          `json:"id"`
	Email               string             `json:"email"`
	DisplayName         string             `json:"displayName"`
	IsActive            bool               `json:"isActive"`
	FirstName           string             `json:"firstName"`
	Lastname            string             `json:"lastName"`
	Address             *string            `json:"address"`
	PhotoURL            *string            `json:"photoUrl"`
	IsMfaEmailEnabled   bool               `json:"isMfaEmailEnabled"`
	IsMfaAuthAppEnabled bool               `json:"isMfaAuthAppEnabled"`
	IsEmailVerified     bool               `json:"isEmailVerified"`
	Badges              []UserRoleResponse `json:"badges"`
}

type UserRoleResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
