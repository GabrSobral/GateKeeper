package resetpassword

import (
	"context"
	"strings"
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
	UserRepository          repository_interfaces.IUserRepository
	RefreshTokenRepository  repository_interfaces.IRefreshTokenRepository
	PasswordResetRepository repository_interfaces.IPasswordResetRepository
}

type Request struct {
	PasswordResetToken string    `json:"password_reset_token" validate:"required"`
	PasswordResetId    uuid.UUID `json:"password_reset_id" validate:"required"`
	NewPassword        string    `json:"new_password" validate:"required"`
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &ResetPasswordService{
		UserRepository:          repository_handlers.UserRepository{Store: q},
		RefreshTokenRepository:  repository_handlers.RefreshTokenRepository{Store: q},
		PasswordResetRepository: repository_handlers.PasswordResetRepository{Store: q},
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

	if strings.Trim(passwordResetToken.Token, " ") != strings.Trim(request.PasswordResetToken, " ") {
		return &errors.ErrPasswordResetTokenMismatch
	}

	user, err := fp.UserRepository.GetUserByID(ctx, passwordResetToken.UserID)

	if err != nil {
		return err
	}

	if user == nil {
		return &errors.ErrUserNotFound
	}

	hashedPassword, err := application_utils.HashPassword(request.NewPassword)

	if err != nil {
		return err
	}

	user.PasswordHash = &hashedPassword

	fp.UserRepository.UpdateUser(ctx, user)
	fp.RefreshTokenRepository.RevokeRefreshTokenFromUser(ctx, user.ID)
	fp.PasswordResetRepository.DeletePasswordResetFromUser(ctx, user.ID)

	return nil
}
