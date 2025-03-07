package session

import (
	"context"

	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/google/uuid"
)

type SessionService struct{}

type Request struct {
	AccessToken string
}

type Response struct {
	User        UserData `json:"user"`
	AccessToken string   `json:"accessToken"`
}

type UserData struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"displayName"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhotoURL    *string   `json:"photoUrl"`
}

func New() *SessionService {
	return &SessionService{}
}

func (ss *SessionService) Handler(ctx context.Context, request Request) (*Response, error) {
	token, err := application_utils.DecodeToken(request.AccessToken)

	if err != nil {
		return nil, err
	}

	return &Response{
		User: UserData{
			ID:          token.UserID,
			DisplayName: token.DisplayName,
			FirstName:   token.FirstName,
			LastName:    token.LastName,
			Email:       token.Email,
			PhotoURL:    nil,
		},
		AccessToken: request.AccessToken,
	}, nil
}
