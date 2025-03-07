package resendemailconfirmation

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	mailservice "github.com/gate-keeper/internal/infra/mail-service"
	"github.com/google/uuid"
)

type Request struct {
	ApplicationID uuid.UUID `json:"applicationId"`
	Email         string    `json:"email"`
}

type ResendEmailConfirmation struct {
	ApplicationUserRepository   repository_interfaces.IApplicationUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository

	MailService mailservice.IMailService
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &ResendEmailConfirmation{
		ApplicationUserRepository:   repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:       repository_handlers.UserProfileRepository{Store: q},
		EmailConfirmationRepository: repository_handlers.EmailConfirmationRepository{Store: q},
		MailService:                 &mailservice.MailService{},
	}
}

func (cm *ResendEmailConfirmation) Handler(ctx context.Context, request Request) error {
	user, err := cm.ApplicationUserRepository.GetUserByEmail(ctx, request.Email, request.ApplicationID)

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

	if emailConfirmation != nil && emailConfirmation.CoolDown.After(time.Now().UTC()) {
		return &errors.ErrEmailConfirmationIsInCoolDown
	}

	if emailConfirmation != nil {
		cm.EmailConfirmationRepository.DeleteEmailConfirmation(ctx, emailConfirmation.ID)
	}

	expiresAt := time.Now().UTC().Add(20 * time.Minute) // 20 minutes
	newEmailConfirmation := entities.NewEmailConfirmation(user.ID, user.Email, expiresAt)

	if err := cm.EmailConfirmationRepository.AddEmailConfirmation(ctx, newEmailConfirmation); err != nil {
		return err
	}

	userProfile, err := cm.UserProfileRepository.GetUserById(ctx, user.ID)

	if err != nil {
		return err
	}

	if err := cm.MailService.SendEmailConfirmationEmail(ctx, user.Email, userProfile.FirstName, newEmailConfirmation.Token); err != nil {
		panic(err)
	}

	return nil
}
