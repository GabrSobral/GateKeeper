package confirmuseremail

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type Handler struct {
	repository IRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Command, *Response] {
	return &Handler{
		repository: Repository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, command Command) (*Response, error) {
	user, err := s.repository.GetUserByEmail(ctx, command.Email, command.ApplicationID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	authorizationCode, err := entities.CreateApplicationAuthorizationCode(
		command.ApplicationID,
		user.ID,
		command.RedirectUri,
		command.CodeChallenge,
		command.CodeChallengeMethod,
	)

	if err != nil {
		return nil, err
	}

	if err := s.repository.AddAuthorizationCode(ctx, authorizationCode); err != nil {
		return nil, err
	}

	emailConfirmation, err := s.repository.GetEmailConfirmationByEmail(ctx, command.Email, user.ID)

	if err != nil {
		return nil, nil
	}

	if emailConfirmation == nil {
		return nil, &errors.ErrEmailConfirmationNotFound
	}

	if emailConfirmation.Token != command.Token {
		return nil, &errors.ErrConfirmationTokenInvalid
	}

	if emailConfirmation.IsUsed {
		return nil, &errors.ErrConfirmationTokenAlreadyUsed
	}

	if emailConfirmation.ExpiresAt.Before(time.Now().UTC()) {
		return nil, &errors.ErrConfirmationTokenAlreadyExpired
	}

	user.IsEmailConfirmed = true
	emailConfirmation.IsUsed = true

	if _, err := s.repository.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	if err := s.repository.UpdateEmailConfirmation(ctx, emailConfirmation); err != nil {
		return nil, err
	}

	return &Response{
		AuthorizationCode:   authorizationCode.ID.String(),
		RedirectUri:         command.RedirectUri,
		State:               command.State,
		CodeChallenge:       command.CodeChallenge,
		CodeChallengeMethod: command.CodeChallengeMethod,
		ResponseType:        command.ResponseType,
	}, nil
}
