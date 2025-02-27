package createapplicationuser

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type RequestBody struct {
	DisplayName           string      `json:"displayName" validate:"required,min=3,max=100"`
	FirstName             string      `json:"firstName" validate:"required,min=3,max=100"`
	LastName              string      `json:"lastName" validate:"required,min=3,max=100"`
	Email                 string      `json:"email" validate:"required,email"`
	IsEmailConfirmed      bool        `json:"isEmailConfirmed" validate:"required"`
	TemporaryPasswordHash *string     `json:"temporaryPasswordHash" validate:"min=8,max=100"`
	IsMfaAuthAppEnabled   bool        `json:"isMfaAuthAppEnabled" validate:"boolean"`
	IsMfaEmailEnabled     bool        `json:"isMfaEmailEnabled" validate:"boolean"`
	Roles                 []uuid.UUID `json:"roles" validate:"required"`
}

type Request struct {
	ApplicationID         uuid.UUID
	DisplayName           string
	FirstName             string
	LastName              string
	Email                 string
	IsEmailConfirmed      bool
	TemporaryPasswordHash *string
	IsMfaAuthAppEnabled   bool
	IsMfaEmailEnabled     bool
	Roles                 []uuid.UUID
}

type Response struct {
	ID          uuid.UUID          `json:"id"`
	DisplayName string             `json:"displayName"`
	Email       string             `json:"email"`
	Roles       []applicationRoles `json:"roles"`
}

type applicationRoles struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

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
	isApplicationExists, err := s.ApplicationRepository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if !isApplicationExists {
		return nil, &errors.ErrApplicationNotFound
	}

	isEmailExists, err := s.ApplicationUserRepository.IsUserExistsByEmail(ctx, request.Email, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if isEmailExists {
		return nil, &errors.ErrUserAlreadyExists
	}

	applicationUser, err := entities.CreateApplicationUser(
		request.Email,
		request.TemporaryPasswordHash,
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

	err = s.ApplicationUserRepository.AddUser(ctx, applicationUser)

	if err != nil {
		return nil, err
	}

	err = s.UserProfileRepository.AddUserProfile(ctx, applicationUserProfile)

	if err != nil {
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
