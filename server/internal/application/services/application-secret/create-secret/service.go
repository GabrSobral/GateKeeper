package createsecret

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

type ControllerRequest struct {
	Name      string     `json:"name" validate:"required"`
	ExpiresAt *time.Time `json:"expiresAt"`
}

type Request struct {
	ApplicationID uuid.UUID  `json:"applicationId" validate:"required,uuid"`
	Name          string     `json:"name" validate:"required"`
	ExpiresAt     *time.Time `json:"expiresAt"`
}

type Response struct {
	ID            uuid.UUID  `json:"id"`
	ApplicationID uuid.UUID  `json:"applicationId"`
	Name          string     `json:"name"`
	Value         string     `json:"value"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	ExpiresAt     *time.Time `json:"expiresAt"`
}

type CreateApplicationSecretService struct {
	ApplicationRepository       repository_interfaces.IApplicationRepository
	ApplicationSecretRepository repository_interfaces.IApplicationSecretRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &CreateApplicationSecretService{
		ApplicationRepository:       repository_handlers.ApplicationRepository{Store: q},
		ApplicationSecretRepository: repository_handlers.ApplicationSecretRepository{Store: q},
	}
}

func (s *CreateApplicationSecretService) Handler(ctx context.Context, request Request) (*Response, error) {
	isApplicationExists, err := s.ApplicationRepository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return nil, err
	}

	if !isApplicationExists {
		return nil, &errors.ErrAplicationNotFound
	}

	newSecret := entities.NewApplicationSecret(request.ApplicationID, request.Name, request.ExpiresAt)

	if err := s.ApplicationSecretRepository.AddSecret(ctx, newSecret); err != nil {
		return nil, err
	}

	return &Response{
		ID:            newSecret.ID,
		ApplicationID: newSecret.ApplicationID,
		Name:          newSecret.Name,
		Value:         newSecret.Value,
		CreatedAt:     newSecret.CreatedAt,
		UpdatedAt:     newSecret.UpdatedAt,
		ExpiresAt:     newSecret.ExpiresAt,
	}, nil
}
