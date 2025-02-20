package http_controllers

import (
	"net/http"

	createapplication "github.com/gate-keeper/internal/application/services/application/create-application"
	getapplicationbyid "github.com/gate-keeper/internal/application/services/application/get-application-by-id"
	listapplications "github.com/gate-keeper/internal/application/services/application/list-applications"
	removeapplication "github.com/gate-keeper/internal/application/services/application/remove-application"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationController struct {
	DbPool *pgxpool.Pool
}

func (c *ApplicationController) GetApplicationByID(writter http.ResponseWriter, request *http.Request) {
	organizationIDString := chi.URLParam(request, "organizationID")
	organizationIdUUID, err := uuid.Parse(organizationIDString)

	applicationIDString := chi.URLParam(request, "applicationID")
	applicationnIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	requestSchema := getapplicationbyid.Request{
		OrganizationID: organizationIdUUID,
		ApplicationID:  applicationnIdUUID,
	}

	params := repositories.ParamsRs[getapplicationbyid.Request, *getapplicationbyid.Response, getapplicationbyid.GetApplicationByIDService]{
		DbPool:  c.DbPool,
		New:     getapplicationbyid.New,
		Request: requestSchema,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

func (c *ApplicationController) ListApplications(writter http.ResponseWriter, request *http.Request) {
	organizationIDString := chi.URLParam(request, "organizationID")
	organizationIdUUID, err := uuid.Parse(organizationIDString)

	if err != nil {
		panic(err)
	}

	requestSchema := listapplications.Request{OrganizationID: organizationIdUUID}

	params := repositories.ParamsRs[listapplications.Request, *[]listapplications.Response, listapplications.ListApplicationsService]{
		DbPool:  c.DbPool,
		New:     listapplications.New,
		Request: requestSchema,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}

func (c *ApplicationController) CreateApplication(writter http.ResponseWriter, request *http.Request) {
	var createApplicationRequest createapplication.Request

	if err := http_router.ParseBodyToSchema(&createApplicationRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[createapplication.Request, *createapplication.Response, createapplication.CreateApplicationService]{
		DbPool:  c.DbPool,
		New:     createapplication.New,
		Request: createApplicationRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusCreated)
}

func (c *ApplicationController) RemoveApplication(writter http.ResponseWriter, request *http.Request) {
	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	organizationIDString := chi.URLParam(request, "organizationID")
	organizationIdUUID, err := uuid.Parse(organizationIDString)

	if err != nil {
		panic(err)
	}

	requestSchema := removeapplication.Request{
		ApplicationID:  applicationIdUUID,
		OrganizationID: organizationIdUUID,
	}

	requestSchema.ApplicationID = applicationIdUUID

	params := repositories.Params[removeapplication.Request, removeapplication.RemoveApplicationService]{
		DbPool:  c.DbPool,
		New:     removeapplication.New,
		Request: requestSchema,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}
