package changepassword

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

type ChangePasswordService struct {
	ApplicationRepository        repository_interfaces.IApplicationRepository
	ApplicationUserRepository    repository_interfaces.IApplicationUserRepository
	RefreshTokenRepository       repository_interfaces.IRefreshTokenRepository
	ChangePasswordCodeRepository repository_interfaces.IChangePasswordCodeRepository
}

type Request struct {
	ChangePasswordCode string    `json:"changePasswordCode" validate:"required,min=64,max=64"`
	ApplicationID      uuid.UUID `json:"applicationID" validate:"required"`
	UserID             uuid.UUID `json:"userID" validate:"required"`
	NewPassword        string    `json:"newPassword" validate:"required,min=8,max=64"`
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &ChangePasswordService{
		ApplicationRepository:        repository_handlers.ApplicationRepository{Store: q},
		ApplicationUserRepository:    repository_handlers.ApplicationUserRepository{Store: q},
		RefreshTokenRepository:       repository_handlers.RefreshTokenRepository{Store: q},
		ChangePasswordCodeRepository: repository_handlers.ChangePasswordCodeRepository{Store: q},
	}
}

func (ss *ChangePasswordService) Handler(ctx context.Context, request Request) error {
	changePasswordCode, err := ss.ChangePasswordCodeRepository.GetByToken(ctx, request.UserID, request.ChangePasswordCode)

	if err != nil {
		return err
	}

	if changePasswordCode == nil {
		return &errors.ErrChangePasswordCodeNotFound
	}

	if changePasswordCode.ExpiresAt.Before(time.Now().UTC()) {
		return &errors.ErrChangePasswordCodeExpired
	}

	if changePasswordCode.Token != request.ChangePasswordCode {
		return &errors.ErrChangePasswordTokenMismatch
	}

	user, err := ss.ApplicationUserRepository.GetUserByID(ctx, changePasswordCode.UserID)

	if err != nil {
		return err
	}

	if user == nil {
		return &errors.ErrUserNotFound
	}

	if user.ShouldChangePass == false {
		return &errors.ErrUserShouldNotChangePassword
	}

	application, err := ss.ApplicationRepository.GetApplicationByID(ctx, request.ApplicationID)

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
	user.ShouldChangePass = false

	ss.ApplicationUserRepository.UpdateUser(ctx, user)
	ss.RefreshTokenRepository.RevokeRefreshTokenFromUser(ctx, user.ID)
	ss.ChangePasswordCodeRepository.RevokeAllByID(ctx, user.ID)

	return nil
}
