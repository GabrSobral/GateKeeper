package editapplicationuser

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
)

type CreateApplicationUserService struct {
	ApplicationRepository      repository_interfaces.IApplicationRepository
	ApplicationUserRepository  repository_interfaces.IApplicationUserRepository
	UserRoleRepository         repository_interfaces.IUserRoleRepository
	UserProfileRepository      repository_interfaces.IUserProfileRepository
	ApplicationRolesRepository repository_interfaces.IApplicationRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &CreateApplicationUserService{
		ApplicationRepository:      repository_handlers.ApplicationRepository{Store: q},
		ApplicationUserRepository:  repository_handlers.ApplicationUserRepository{Store: q},
		UserRoleRepository:         repository_handlers.UserRoleRepository{Store: q},
		UserProfileRepository:      repository_handlers.UserProfileRepository{Store: q},
		ApplicationRolesRepository: repository_handlers.ApplicationRoleRepository{Store: q},
	}
}

func (s *CreateApplicationUserService) Handler(ctx context.Context, request Request) (*Response, error) {
	application, err := s.ApplicationRepository.GetApplicationByID(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if application == nil {
		return nil, &errors.ErrApplicationNotFound
	}

	applicationUser, err := s.ApplicationUserRepository.GetUserByID(ctx, request.UserID)

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

	if _, err = s.ApplicationUserRepository.UpdateUser(ctx, applicationUser); err != nil {
		return nil, err
	}

	if err = s.UserProfileRepository.EditUserProfile(ctx, applicationUserProfile); err != nil {
		return nil, err
	}

	userRoles, err := s.UserRoleRepository.GetRolesByUserID(ctx, applicationUser.ID)

	if err != nil {
		return nil, err
	}

	for _, role := range userRoles {
		userRole := entities.UserRole{
			UserID: applicationUser.ID,
			RoleID: role.ID,
		}

		err = s.UserRoleRepository.RemoveUserRole(ctx, &userRole)
	}

	for _, roleID := range request.Roles {
		userRole := entities.UserRole{
			UserID:    applicationUser.ID,
			CreatedAt: time.Now(),
			RoleID:    roleID,
		}

		err = s.UserRoleRepository.AddUserRole(ctx, &userRole)
	}

	roles := make([]applicationRoles, len(request.Roles))
	applicationRolesList, err := s.ApplicationRolesRepository.ListRolesFromApplication(ctx, request.ApplicationID)

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
