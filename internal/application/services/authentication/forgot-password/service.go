package forgotpassword

import (
	"context"

	"github.com/guard-service/internal/domain/entities"
	"github.com/guard-service/internal/domain/errors"
	repository_interfaces "github.com/guard-service/internal/infra/database/repositories/interfaces"
	mailservice "github.com/guard-service/internal/infra/mail-service"
)

type ForgotPasswordService struct {
	PasswordResetRepository repository_interfaces.IPasswordResetRepository
	UserRepository          repository_interfaces.IUserRepository
	UserProfileRepository   repository_interfaces.IUserProfileRepository
	MailService             mailservice.IMailService
}

type Request struct {
	Email string `json:"email" validate:"required,email"`
}

func (fp *ForgotPasswordService) Handler(ctx context.Context, request Request) error {
	user, err := fp.UserRepository.GetUserByEmail(ctx, request.Email)

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
		if err := fp.MailService.SendForgotPasswordEmail(ctx, user.Email, userProfile.FirstName, passwordResetToken.Token, passwordResetToken.ID); err != nil {
			panic(err)
		}
	}()

	return nil
}
