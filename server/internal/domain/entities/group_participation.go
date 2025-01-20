package entities

import (
	"time"

	"github.com/google/uuid"
)

type GroupParticipation struct {
	GroupID   uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewGroupParticipation(userID uuid.UUID, groupID uuid.UUID) *GroupParticipation {
	return &GroupParticipation{
		UserID:    userID,
		GroupID:   groupID,
		CreatedAt: time.Now(),
	}
}
