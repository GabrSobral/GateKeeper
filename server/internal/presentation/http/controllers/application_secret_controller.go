package http_controllers

import (
	"net/http"

	createsecret "github.com/gate-keeper/internal/application/services/application-secret/create-secret"
	deletesecret "github.com/gate-keeper/internal/application/services/application-secret/delete-secret"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationSecretController struct {
	DbPool *pgxpool.Pool
}

func (c *ApplicationSecretController) CreateSecret(writter http.ResponseWriter, request *http.Request) {
	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	var controllerRequest createsecret.ControllerRequest

	if err := http_router.ParseBodyToSchema(&controllerRequest, request); err != nil {
		panic(err)
	}

	schema := createsecret.Request{
		ApplicationID: applicationIdUUID,
		Name:          controllerRequest.Name,
		ExpiresAt:     controllerRequest.ExpiresAt,
	}

	params := repositories.ParamsRs[createsecret.Request, *createsecret.Response, createsecret.CreateApplicationSecretService]{
		DbPool:  c.DbPool,
		New:     createsecret.New,
		Request: schema,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusCreated)
}

func (c *ApplicationSecretController) RemoveSecret(writter http.ResponseWriter, request *http.Request) {
	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	secretIDString := chi.URLParam(request, "secretID")
	secretIdUUID, err := uuid.Parse(secretIDString)

	if err != nil {
		panic(err)
	}

	schema := deletesecret.Request{
		SecretID:      secretIdUUID,
		ApplicationID: applicationIdUUID,
	}

	params := repositories.Params[deletesecret.Request, deletesecret.DeleteSecretService]{
		DbPool:  c.DbPool,
		New:     deletesecret.New,
		Request: schema,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}
