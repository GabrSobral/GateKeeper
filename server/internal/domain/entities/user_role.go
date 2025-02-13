package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	UserID    uuid.UUID
	RoleID    uuid.UUID
	CreatedAt time.Time
}

func NewUserRole(userID, roleID uuid.UUID) *UserRole {
	return &UserRole{
		UserID:    userID,
		RoleID:    roleID,
		CreatedAt: time.Now(),
	}
}
