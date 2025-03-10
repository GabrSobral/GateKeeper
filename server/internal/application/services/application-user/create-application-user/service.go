package createapplicationuser

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
	ApplicationRepository     repository_interfaces.IApplicationRepository
	ApplicationUserRepository repository_interfaces.IApplicationUserRepository
	UserRoleRepository        repository_interfaces.IUserRoleRepository
	UserProfileRepository     repository_interfaces.IUserProfileRepository
	ApplicationRoles          repository_interfaces.IApplicationRoleRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &CreateApplicationUserService{
		ApplicationRepository:     repository_handlers.ApplicationRepository{Store: q},
		ApplicationUserRepository: repository_handlers.ApplicationUserRepository{Store: q},
		UserRoleRepository:        repository_handlers.UserRoleRepository{Store: q},
		UserProfileRepository:     repository_handlers.UserProfileRepository{Store: q},
		ApplicationRoles:          repository_handlers.ApplicationRoleRepository{Store: q},
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

	isEmailExists, err := s.ApplicationUserRepository.IsUserExistsByEmail(ctx, request.Email, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if isEmailExists {
		return nil, &errors.ErrUserAlreadyExists
	}

	hashedPassword, err := application_utils.HashPassword(*request.TemporaryPasswordHash, application.PasswordHashSecret)

	if err != nil {
		return nil, err
	}

	applicationUser, err := entities.CreateApplicationUser(
		request.Email,
		&hashedPassword,
		request.ApplicationID,
		true, // shouldChangePass
	)

	if err != nil {
		return nil, err
	}

	applicationUser.IsMfaAuthAppEnabled = request.IsMfaAuthAppEnabled
	applicationUser.IsMfaEmailEnabled = request.IsMfaEmailEnabled

	applicationUserProfile := entities.NewUserProfile(
		applicationUser.ID,
		request.FirstName,
		request.LastName,
		request.DisplayName,
		nil,
		nil,
		nil,
	)

	if err = s.ApplicationUserRepository.AddUser(ctx, applicationUser); err != nil {
		return nil, err
	}

	if err = s.UserProfileRepository.AddUserProfile(ctx, applicationUserProfile); err != nil {
		return nil, err
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
	applicationRolesList, err := s.ApplicationRoles.ListRolesFromApplication(ctx, request.ApplicationID)

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
