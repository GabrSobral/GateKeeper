package http_controllers

import (
	"net/http"

	createapplication "github.com/gate-keeper/internal/application/services/application/create-application"
	removeapplication "github.com/gate-keeper/internal/application/services/application/remove-application"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationController struct {
	DbPool *pgxpool.Pool
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

	var removeApplicationRequest removeapplication.Request

	if err := http_router.ParseBodyToSchema(&removeApplicationRequest, request); err != nil {
		panic(err)
	}

	removeApplicationRequest.ApplicationID = applicationIdUUID

	params := repositories.Params[removeapplication.Request, removeapplication.RemoveApplicationService]{
		DbPool:  c.DbPool,
		New:     removeapplication.New,
		Request: removeApplicationRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}
