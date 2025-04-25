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

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Request struct {
	ApplicationID uuid.UUID `json:"applicationId" validate:"required"`
	UserID        uuid.UUID `json:"userId" validate:"required"`
}

type Response struct {
	OtpUrl string `json:"otpUrl"`
}

type GenerateAuthAppSecretService struct {
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	ApplicationRepository     repository_interfaces.IApplicationRepository

	MailService mailservice.IMailService
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &GenerateAuthAppSecretService{
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},

		MailService: &mailservice.MailService{},
	}
}

func (ss *GenerateAuthAppSecretService) Handler(ctx context.Context, request Request) (*Response, error) {
	application, err := ss.ApplicationRepository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if application == nil {
		return nil, &errors.ErrApplicationNotFound
	}

	user, err := ss.ApplicationUserRepository.GetUserByID(ctx, request.UserID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	if !user.IsMfaAuthAppEnabled {
		return nil, &errors.ErrMfaAuthAppNotEnabled
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      application.Name,
		AccountName: user.Email,
		Period:      30,
		SecretSize:  32,
		Secret:      []byte(entities.GenerateRandomString(16)),
		Digits:      6,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        nil,
	})

	if err != nil {
		return nil, err
	}

	secret := key.Secret()

	user.TwoFactorSecret = &secret

	ss.ApplicationUserRepository.UpdateUser(ctx, user)

	return &Response{
		OtpUrl: key.URL(),
	}, nil
}
