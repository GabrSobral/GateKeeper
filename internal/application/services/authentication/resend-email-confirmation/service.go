package resendemailconfirmation

import (
	"context"
	"time"

	"github.com/guard-service/internal/domain/entities"
	"github.com/guard-service/internal/domain/errors"
	repository_interfaces "github.com/guard-service/internal/infra/database/repositories/interfaces"
	mailservice "github.com/guard-service/internal/infra/mail-service"
)

type Request struct {
	Email string `json:"email"`
}

type ResendEmailConfirmation struct {
	UserRepository              repository_interfaces.IUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository

	MailService mailservice.IMailService
}

func (cm *ResendEmailConfirmation) Handler(ctx context.Context, request Request) error {
	user, err := cm.UserRepository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil
	}

	if user == nil {
		return &errors.ErrUserNotFound
	}

	emailConfirmation, err := cm.EmailConfirmationRepository.GetByEmail(ctx, request.Email, user.ID)

	if err != nil {
		return err
	}

	if emailConfirmation.CoolDown.After(time.Now().UTC()) {
		return &errors.ErrEmailConfirmationIsInCoolDown
	}

	userProfile, err := cm.UserProfileRepository.GetUserById(ctx, user.ID)

	if err != nil {
		return err
	}

	cm.EmailConfirmationRepository.DeleteEmailConfirmation(ctx, emailConfirmation.ID)

	expiresAt := time.Now().UTC().Add(20 * time.Minute)
	newEmailConfirmation := entities.NewEmailConfirmation(user.ID, user.Email, expiresAt)

	if err := cm.EmailConfirmationRepository.AddEmailConfirmation(ctx, newEmailConfirmation); err != nil {
		return err
	}

	if err := cm.MailService.SendEmailConfirmationEmail(ctx, user.Email, userProfile.FirstName, newEmailConfirmation.Token); err != nil {
		panic(err)
	}

	return nil
}
