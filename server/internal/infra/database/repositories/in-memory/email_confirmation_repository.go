package inmemory_repositories

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type InMemoryEmailConfirmationRepository struct {
	Emails map[string]*entities.EmailConfirmation
}

func (r InMemoryEmailConfirmationRepository) AddEmailConfirmation(ctx context.Context, emailConfirmation *entities.EmailConfirmation) error {
	if r.Emails == nil {
		r.Emails = make(map[string]*entities.EmailConfirmation)
	}

	r.Emails[emailConfirmation.ID.String()] = emailConfirmation

	return nil
}

func (r InMemoryEmailConfirmationRepository) GetByEmail(ctx context.Context, email string, userID uuid.UUID) (*entities.EmailConfirmation, error) {
	for _, emailConfirmation := range r.Emails {
		if emailConfirmation.Email == email && emailConfirmation.UserID == userID {
			return emailConfirmation, nil
		}
	}

	return nil, nil
}

func (r InMemoryEmailConfirmationRepository) UpdateEmailConfirmation(ctx context.Context, emailConfirmation *entities.EmailConfirmation) error {
	r.Emails[emailConfirmation.ID.String()] = emailConfirmation

	return nil
}

func (r InMemoryEmailConfirmationRepository) DeleteEmailConfirmation(ctx context.Context, emailConfirmationID uuid.UUID) error {
	delete(r.Emails, emailConfirmationID.String())

	return nil
}
