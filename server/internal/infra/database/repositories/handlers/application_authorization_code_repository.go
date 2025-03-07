package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationAuthorizationCodeRepository struct {
	Store *pgstore.Queries
}

func (r ApplicationAuthorizationCodeRepository) AddAuthorizationCode(ctx context.Context, newAuthorizationCode *entities.ApplicationAuthorizationCode) error {
	err := r.Store.AddAuthorizationCode(ctx, pgstore.AddAuthorizationCodeParams{
		ID:                  newAuthorizationCode.ID,
		ApplicationID:       newAuthorizationCode.ApplicationID,
		UserID:              newAuthorizationCode.ApplicationUserId,
		ExpiredAt:           pgtype.Timestamp{Time: newAuthorizationCode.ExpiresAt, Valid: true},
		Code:                newAuthorizationCode.Code,
		RedirectUri:         newAuthorizationCode.RedirectUri,
		CodeChallenge:       newAuthorizationCode.CodeChallenge,
		CodeChallengeMethod: newAuthorizationCode.CodeChallengeMethod,
	})

	return err
}

func (r ApplicationAuthorizationCodeRepository) RemoveAuthorizationCode(ctx context.Context, userID, applicationId uuid.UUID) error {
	err := r.Store.RemoveAuthorizationCode(ctx, pgstore.RemoveAuthorizationCodeParams{
		ApplicationID: applicationId,
		UserID:        userID,
	})

	return err
}

func (r ApplicationAuthorizationCodeRepository) GetAuthorizationCodeById(ctx context.Context, code uuid.UUID) (*entities.ApplicationAuthorizationCode, error) {
	authorizationCode, err := r.Store.GetAuthorizationCodeById(ctx, code)

	if err != nil && err != repositories.ErrNoRows {
		return nil, err
	}

	return &entities.ApplicationAuthorizationCode{
		ID:                  authorizationCode.ID,
		ApplicationID:       authorizationCode.ApplicationID,
		ExpiresAt:           authorizationCode.ExpiredAt.Time,
		Code:                authorizationCode.Code,
		ApplicationUserId:   authorizationCode.UserID,
		RedirectUri:         authorizationCode.RedirectUri,
		CodeChallenge:       authorizationCode.CodeChallenge,
		CodeChallengeMethod: authorizationCode.CodeChallengeMethod,
	}, nil
}
