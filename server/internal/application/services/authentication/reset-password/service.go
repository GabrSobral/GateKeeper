package resetpassword

import (
	"context"
	"time"

	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type ResetPasswordService struct {
	ApplicationRepository     repository_interfaces.IApplicationRepository
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	RefreshTokenRepository    repository_interfaces.IRefreshTokenRepository
	PasswordResetRepository   repository_interfaces.IPasswordResetRepository
}

type Request struct {
	PasswordResetToken string    `json:"passwordResetToken" validate:"required"`
	PasswordResetId    uuid.UUID `json:"passwordResetId" validate:"required"`
	NewPassword        string    `json:"newPassword" validate:"required"`
	ApplicationID      uuid.UUID `json:"applicationId" validate:"required"`
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &ResetPasswordService{
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		RefreshTokenRepository:    repository_handlers.RefreshTokenRepository{Store: q},
		PasswordResetRepository:   repository_handlers.PasswordResetRepository{Store: q},
	}
}

func (fp *ResetPasswordService) Handler(ctx context.Context, request Request) error {
	passwordResetToken, err := fp.PasswordResetRepository.GetByTokenID(ctx, request.PasswordResetId)

	if err != nil {
		return err
	}

	if passwordResetToken == nil {
		return &errors.ErrPasswordResetNotFound
	}

	if passwordResetToken.ExpiresAt.Before(time.Now().UTC()) {
		return &errors.ErrPasswordResetTokenExpired
	}

	if passwordResetToken.Token != request.PasswordResetToken {
		return &errors.ErrPasswordResetTokenMismatch
	}

	user, err := fp.ApplicationUserRepository.GetUserByID(ctx, passwordResetToken.UserID)

	if err != nil {
		return err
	}

	if user == nil {
		return &errors.ErrUserNotFound
	}

	application, err := fp.ApplicationRepository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	if application == nil {
		return &errors.ErrApplicationNotFound
	}

	hashedPassword, err := application_utils.HashPassword(request.NewPassword, application.PasswordHashSecret)

	if err != nil {
		return err
	}

	user.PasswordHash = &hashedPassword

	fp.ApplicationUserRepository.UpdateUser(ctx, user)
	fp.RefreshTokenRepository.RevokeRefreshTokenFromUser(ctx, user.ID)
	fp.PasswordResetRepository.DeletePasswordResetFromUser(ctx, user.ID)

	return nil
}
