package resendemailconfirmation

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	mailservice "github.com/gate-keeper/internal/infra/mail-service"
	"github.com/google/uuid"
)

type ForgotPasswordService struct {
	PasswordResetRepository   repository_interfaces.IPasswordResetRepository
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	UserProfileRepository     repository_interfaces.IUserProfileRepository
	MailService               mailservice.IMailService
}

type Request struct {
	ApplicationID uuid.UUID `json:"applicationId" validate:"required,uuid"`
	Email         string    `json:"email" validate:"required,email"`
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &ForgotPasswordService{
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
		PasswordResetRepository:   repository_handlers.PasswordResetRepository{Store: q},
		MailService:               &mailservice.MailService{},
	}
}

func (fp *ForgotPasswordService) Handler(ctx context.Context, request Request) error {
	user, err := fp.ApplicationUserRepository.GetUserByEmail(ctx, request.Email, request.ApplicationID)

	if err != nil {
		return err
	}

	if user == nil {
		return &errors.ErrUserNotFound
	}

	if !user.IsEmailConfirmed {
		return &errors.ErrEmailNotConfirmed
	}

	fp.PasswordResetRepository.DeletePasswordResetFromUser(ctx, user.ID)

	passwordResetToken, err := entities.NewPasswordResetToken(user.ID)

	if err != nil {
		return err
	}

	if err := fp.PasswordResetRepository.CreatePasswordReset(ctx, passwordResetToken); err != nil {
		return nil
	}

	userProfile, err := fp.UserProfileRepository.GetUserById(ctx, user.ID)

	if err != nil {
		return nil
	}

	go func() {
		if err := fp.MailService.SendForgotPasswordEmail(ctx, user.Email, userProfile.FirstName, passwordResetToken.Token, passwordResetToken.ID, request.ApplicationID); err != nil {
			panic(err)
		}
	}()

	return nil
}
