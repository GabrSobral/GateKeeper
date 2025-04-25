package getapplicationauthdata

import (
	"context"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type Handler struct {
	repository IRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Query, *Response] {
	return &Handler{
		repository: Repository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, query Query) (*Response, error) {
	application, err := s.repository.GetApplicationByID(ctx, query.ApplicationID)

	if err != nil {
		return nil, err
	}

	if application == nil {
		return nil, &errors.ErrApplicationNotFound
	}

	return &Response{
		ID:                application.ID,
		Name:              application.Name,
		CanSelfSignUp:     application.CanSelfSignUp,
		CanSelfForgotPass: application.CanSelfForgotPass,
		OAuthProviders:    make([]ApplicationProviders, 0),
	}, nil
}
