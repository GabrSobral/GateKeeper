package configureoauthprovider

import (
	"context"
	"fmt"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type Handler struct {
	repository IRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandler[Command] {
	return &Handler{
		repository: Repository{Store: q},
	}
}

func (s *Handler) Handler(ctx context.Context, request Command) error {
	isApplicationExists, err := s.repository.CheckIfApplicationExists(ctx, request.ApplicationID)

	if err != nil {
		return err
	}

	if !isApplicationExists {
		return &errors.ErrApplicationNotFound
	}

	applicationOauthProvider, err := s.repository.GetApplicationOauthProviderByName(ctx, request.ApplicationID, request.Name)

	fmt.Println("test -2", err)

	if err != nil {
		return err
	}

	fmt.Println("test -1")

	if applicationOauthProvider != nil {
		applicationOauthProvider.ClientID = request.ClientID
		applicationOauthProvider.ClientSecret = request.ClientSecret
		applicationOauthProvider.RedirectURI = request.RedirectURI
		applicationOauthProvider.Enabled = request.Enabled

		err = s.repository.UpdateApplicationOauthProvider(ctx, applicationOauthProvider)

		if err != nil {
			return err
		}

		fmt.Println("test 4")

		return nil
	}

	fmt.Println("test 1")

	newApplicationOauthProvider := entities.NewApplicationOAuthProvider(
		request.ApplicationID,
		request.Name,
		request.ClientID,
		request.ClientSecret,
		request.RedirectURI,
		request.Enabled,
	)

	fmt.Println("test 2")

	err = s.repository.AddApplicationOauthProvider(ctx, newApplicationOauthProvider)

	if err != nil {
		return err
	}

	fmt.Println("test 3")

	return nil
}
