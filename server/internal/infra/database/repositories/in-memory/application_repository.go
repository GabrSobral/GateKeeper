package inmemory_repositories

import "github.com/gate-keeper/internal/domain/entities"

type InMemoryApplicationRepository struct {
	applications map[string]entities.Application
}

func (r InMemoryApplicationRepository) AddApplication(newApplication *entities.Application) error {
	r.applications[newApplication.ID.String()] = *newApplication
	return nil
}

func (r InMemoryApplicationRepository) GetApplicationByID(applicationID string) (*entities.Application, error) {
	application := r.applications[applicationID]
	return &application, nil
}

func (r InMemoryApplicationRepository) RemoveApplication(applicationID string) error {
	delete(r.applications, applicationID)
	return nil
}

func (r InMemoryApplicationRepository) UpdateApplication(applicationID string, newApplication *entities.Application) error {
	r.applications[applicationID] = *newApplication
	return nil
}
