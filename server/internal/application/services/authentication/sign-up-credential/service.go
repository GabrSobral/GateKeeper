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
)

type Request struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type SignUpService struct {
	UserRepository              repository_interfaces.IUserRepository
	UserProfileRepository       repository_interfaces.IUserProfileRepository
	RefreshTokenRepository      repository_interfaces.IRefreshTokenRepository
	EmailConfirmationRepository repository_interfaces.IEmailConfirmationRepository
	MailService                 mailservice.IMailService
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Request] {
	return &SignUpService{
		UserRepository:              repository_handlers.UserRepository{Store: q},
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

	isUserExist, err := ss.UserRepository.IsUserExistsByEmail(ctx, request.Email)

	if err != nil {
		return err
	}

	if isUserExist {
		return &errors.ErrUserAlreadyExists
	}

	hashedPassword, err := application_utils.HashPassword(request.Password)

	if err != nil {
		return err
	}

	user, err := entities.CreateUser(request.Email, &hashedPassword)

	if err != nil {
		return err
	}

	userProfile := entities.NewUserProfile(
		user.ID,
		request.FirstName,
		request.LastName,
		nil,
		nil,
		nil,
	)

	if err := ss.UserRepository.AddUser(ctx, user); err != nil {
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
