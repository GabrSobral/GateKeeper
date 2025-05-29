package login

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	application_utils "github.com/gate-keeper/internal/features/utils"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	mailservice "github.com/gate-keeper/internal/infra/mail-service"
)

type Handler struct {
	repository  IRepository
	mailService mailservice.IMailService
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Command, *Response] {
	return &Handler{
		repository:  Repository{Store: q},
		mailService: &mailservice.MailService{},
	}
}

func (s *Handler) Handler(ctx context.Context, command Command) (*Response, error) {
	user, err := s.repository.GetUserByEmail(ctx, command.Email, command.ApplicationID)

	if err != nil {
		return nil, &errors.ErrUserNotFound
	}

	if user == nil {
		return nil, &errors.ErrUserNotFound
	}

	if !user.IsActive {
		return nil, &errors.ErrUserNotActive
	}

	if user.PasswordHash == nil {
		return nil, &errors.ErrUserSignUpWithSocial
	}

	isPasswordCorrect, err := application_utils.ComparePassword(*user.PasswordHash, command.Password)

	if err != nil {
		return nil, err
	}

	if !isPasswordCorrect {
		return nil, &errors.ErrEmailOrPasswordInvalid
	}

	if !user.IsEmailConfirmed {
		return nil, &errors.ErrEmailNotConfirmed
	}

	// Revoke all Password change codes if exists
	if err := s.repository.RevokeAllChangePasswordCodeByUserID(ctx, user.ID); err != nil {
		return nil, err
	}

	var changePasswordCode *entities.ChangePasswordCode = nil

	if user.ShouldChangePass {
		changePasswordCode = entities.NewChangePasswordCode(user.ID, user.Email)

		if err := s.repository.AddChangePasswordCode(ctx, changePasswordCode); err != nil {
			return nil, err
		}
	}

	if user.Preferred2FAMethod != nil && *user.Preferred2FAMethod == entities.MFAEmail {
		userProfile, err := s.repository.GetUserProfileByID(ctx, user.ID)

		if err != nil {
			panic(err)
		}

		emailMfaCode := entities.NewEmailMfaCode(user.ID, user.Email)

		if err := s.repository.AddEmailMfaCode(ctx, emailMfaCode); err != nil {
			panic(err)
		}

		go func() {
			if err := s.mailService.SendMfaEmail(ctx, user.Email, userProfile.FirstName, emailMfaCode.Token); err != nil {
				panic(err)
			}
		}()

		if changePasswordCode == nil {
			return &Response{
				MfaType:            user.Preferred2FAMethod,
				MfaID:              &emailMfaCode.ID,
				ChangePasswordCode: nil,
				Message:            "MFA is required, please enter the code from your authentication app",
				SessionCode:        nil,
				UserID:             user.ID,
			}, nil
		}

		return &Response{
			MfaType:            user.Preferred2FAMethod,
			MfaID:              &emailMfaCode.ID,
			ChangePasswordCode: &changePasswordCode.Token,
			Message:            "MFA is required, please enter the code from your authentication app",
			SessionCode:        nil,
			UserID:             user.ID,
		}, nil
	}

	if user.Preferred2FAMethod != nil && *user.Preferred2FAMethod == entities.MFAApp {
		appMfaCode := entities.NewAppMfaCode(user.ID, user.Email)

		if err := s.repository.AddAppMfaCode(ctx, appMfaCode); err != nil {
			panic(err)
		}

		if changePasswordCode == nil {
			return &Response{
				MfaType:            user.Preferred2FAMethod,
				MfaID:              &appMfaCode.ID,
				ChangePasswordCode: nil,
				Message:            "MFA is required, please enter the code from your authentication app",
				SessionCode:        nil,
				UserID:             user.ID,
			}, nil
		}

		return &Response{
			MfaType:            user.Preferred2FAMethod,
			MfaID:              &appMfaCode.ID,
			ChangePasswordCode: &changePasswordCode.Token,
			Message:            "MFA is required, please enter the code from your authentication app",
			SessionCode:        nil,
			UserID:             user.ID,
		}, nil
	}

	sessionToken, err := entities.CreateSessionCode(
		user.ID,
		command.ApplicationID,
	)

	if err != nil {
		return nil, err
	}

	if err := s.repository.AddSessionCode(ctx, sessionToken); err != nil {
		return nil, err
	}

	tokenString := sessionToken.Token

	if changePasswordCode == nil {
		return &Response{
			MfaType:            nil,
			MfaID:              nil,
			Message:            "Login successful",
			ChangePasswordCode: nil,
			SessionCode:        &tokenString,
			UserID:             user.ID,
		}, nil
	}

	return &Response{
		MfaType:            nil,
		MfaID:              nil,
		Message:            "Login successful",
		ChangePasswordCode: &changePasswordCode.Token,
		SessionCode:        &tokenString,
		UserID:             user.ID,
	}, nil
}
