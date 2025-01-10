package entities

import "time"

type UserSession struct {
	ID           string
	UserID       string
	CreatedAt    time.Time
	ExpiresAt    time.Time
	DeviceID     string
	LastAccessAt time.Time
	Token        string
	UserAgent    string
	IpAddress    string
}

func NewUserSession(userID, deviceID, userAgent, ipAddress string, expiresAt time.Time) *UserSession {
	return &UserSession{
		UserID:       userID,
		DeviceID:     deviceID,
		UserAgent:    userAgent,
		IpAddress:    ipAddress,
		CreatedAt:    time.Now().UTC(),
		ExpiresAt:    expiresAt,
		LastAccessAt: time.Now().UTC(),
	}
}
