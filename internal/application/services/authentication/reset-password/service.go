package resetpassword

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	application_utils "github.com/guard-service/internal/application/utils"
	"github.com/guard-service/internal/domain/errors"
	repository_interfaces "github.com/guard-service/internal/infra/database/repositories/interfaces"
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
