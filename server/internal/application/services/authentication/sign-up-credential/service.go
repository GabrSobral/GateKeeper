package signup

import (
	"context"
	"time"

	application_utils "github.com/gate-keeper/internal/application/utils"
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
	ApplicationID uuid.UUID `json:"application_id" validate:"required,uuid4"`
	DisplayName   string    `json:"display_name"`
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	Password      string    `json:"password" validate:"required"`
}

type SignUpService struct {
	ApplicationRepository       repository_interfaces.IApplicationRepository
	ApplicationUserRepository   repository_interfaces.IApplicationUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	RefreshTokenRepository      repository_interfaces.IRefreshTokenRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository
	MailService                 mailservice.IMailService
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &SignUpService{
		ApplicationRepository:       repository_handlers.ApplicationRepository{Store: q},
		ApplicationUserRepository:   repository_handlers.ApplicationUserRepository{Store: q},
		UserProfileRepository:       repository_handlers.UserProfileRepository{Store: q},
		RefreshTokenRepository:      repository_handlers.RefreshTokenRepository{Store: q},
		EmailConfirmationRepository: repository_handlers.EmailConfirmationRepository{Store: q},
		MailService:                 &mailservice.MailService{},
	}
}

func (ss *SignUpService) Handler(ctx context.Context, request Request) error {
	isEmailValid := application_utils.EmailValidator(request.Email)

	if !isEmailValid {
		return &errors.ErrInvalidEmail
	}

	isUserExist, err := ss.ApplicationUserRepository.IsUserExistsByEmail(ctx, request.Email, request.ApplicationID)

	if err != nil {
		return err
	}

	if isUserExist {
		return &errors.ErrUserAlreadyExists
	}

	application, err := ss.ApplicationRepository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	hashedPassword, err := application_utils.HashPassword(request.Password, application.PasswordHashSecret)

	if err != nil {
		return err
	}

	user, err := entities.CreateApplicationUser(request.Email, &hashedPassword, request.ApplicationID, false)

	if err != nil {
		return err
	}

	userProfile := entities.NewUserProfile(
		user.ID,
		request.FirstName,
		request.LastName,
		request.DisplayName,
		nil,
		nil,
		nil,
	)

	if err := ss.ApplicationUserRepository.AddUser(ctx, user); err != nil {
		return err
	}

	if err := ss.UserProfileRepository.AddUserProfile(ctx, userProfile); err != nil {
		return err
	}

	expiresAt := time.Now().UTC().Add(20 * time.Minute)
	emailConfirmation := entities.NewEmailConfirmation(user.ID, user.Email, expiresAt)

	if err := ss.EmailConfirmationRepository.AddEmailConfirmation(ctx, emailConfirmation); err != nil {
		return err
	}

	go func() {
		if err := ss.MailService.SendEmailConfirmationEmail(ctx, user.Email, userProfile.FirstName, emailConfirmation.Token); err != nil {
			panic(err)
		}
	}()

	return nil
}
