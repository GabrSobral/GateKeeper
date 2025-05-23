package editapplicationuser

import (
	"context"
	"time"

	application_utils "github.com/gate-keeper/internal/application/utils"
	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type Handler struct {
	repository IRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Command, *Response] {
	return &Handler{
		repository: Repository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, request Command) (*Response, error) {
	application, err := s.repository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if application == nil {
		return nil, &errors.ErrApplicationNotFound
	}

	applicationUser, err := s.repository.GetUserByID(ctx, request.UserID)

	if err != nil {
		return nil, err
	}

	if applicationUser == nil {
		return nil, &errors.ErrUserNotFound
	}

	if request.TemporaryPasswordHash != nil {
		hashedPassword, err := application_utils.HashPassword(*request.TemporaryPasswordHash, application.PasswordHashSecret)

		if err != nil {
			return nil, err
		}

		applicationUser.PasswordHash = &hashedPassword
		applicationUser.ShouldChangePass = true
	}

	currentTime := time.Now().UTC()

	applicationUser.IsMfaAuthAppEnabled = request.IsMfaAuthAppEnabled
	applicationUser.IsMfaEmailEnabled = request.IsMfaEmailEnabled
	applicationUser.UpdatedAt = &currentTime
	applicationUser.IsEmailConfirmed = request.IsEmailConfirmed
	applicationUser.IsActive = request.IsActive

	applicationUserProfile := entities.NewUserProfile(
		applicationUser.ID,
		request.FirstName,
		request.LastName,
		request.DisplayName,
		nil,
		nil,
		nil,
	)

	if _, err = s.repository.UpdateUser(ctx, applicationUser); err != nil {
		return nil, err
	}

	if err = s.repository.EditUserProfile(ctx, applicationUserProfile); err != nil {
		return nil, err
	}

	userRoles, err := s.repository.GetRolesByUserID(ctx, applicationUser.ID)

	if err != nil {
		return nil, err
	}

	for _, role := range userRoles {
		userRole := entities.UserRole{
			UserID: applicationUser.ID,
			RoleID: role.ID,
		}

		err = s.repository.RemoveUserRole(ctx, &userRole)
	}

	for _, roleID := range request.Roles {
		userRole := entities.UserRole{
			UserID:    applicationUser.ID,
			CreatedAt: time.Now(),
			RoleID:    roleID,
		}

		err = s.repository.AddUserRole(ctx, &userRole)
	}

	roles := make([]applicationRoles, len(request.Roles))
	applicationRolesList, err := s.repository.ListRolesFromApplication(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	for i, roleID := range request.Roles {
		for _, appRole := range *applicationRolesList {
			if appRole.ID == roleID {
				roles[i] = applicationRoles{
					ID:          appRole.ID,
					Name:        appRole.Name,
					Description: appRole.Description,
				}
			}
		}
	}

	return &Response{
		ID:          applicationUser.ID,
		DisplayName: applicationUserProfile.DisplayName,
		Email:       applicationUser.Email,
		Roles:       roles,
	}, nil
}
