package http_controllers

import (
	"net/http"

	createorganization "github.com/gate-keeper/internal/application/services/organization/create-organization"
	listorganizations "github.com/gate-keeper/internal/application/services/organization/list-organizations"
	removeorganization "github.com/gate-keeper/internal/application/services/organization/remove-organization"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizationController struct {
	DbPool *pgxpool.Pool
}

func (c *OrganizationController) CreateOrganization(writter http.ResponseWriter, request *http.Request) {
	var createOrganizationRequest createorganization.Request

	if err := http_router.ParseBodyToSchema(&createOrganizationRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[createorganization.Request, *createorganization.Response, createorganization.CreateOrganizationService]{
		DbPool:  c.DbPool,
		New:     createorganization.New,
		Request: createOrganizationRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusCreated)
}

func (c *OrganizationController) RemoveOrganization(writter http.ResponseWriter, request *http.Request) {
	var removeOrganizationRequest removeorganization.Request

	if err := http_router.ParseBodyToSchema(&removeOrganizationRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[removeorganization.Request, removeorganization.RemoveOrganizationService]{
		DbPool:  c.DbPool,
		New:     removeorganization.New,
		Request: removeOrganizationRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusCreated)
}

func (c *OrganizationController) ListOrganizations(writter http.ResponseWriter, request *http.Request) {
	testId, _ := uuid.NewV7()
	var listOrganizationsRequest listorganizations.Request = listorganizations.Request{UserID: testId} // remove later

	params := repositories.ParamsRs[listorganizations.Request, *[]listorganizations.Response, listorganizations.ListOrganizationsService]{
		DbPool:  c.DbPool,
		New:     listorganizations.New,
		Request: listOrganizationsRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}
