package http_controllers

import (
	"net/http"

	createtenant "github.com/gate-keeper/internal/application/services/tenant/create-tenant"
	removetenant "github.com/gate-keeper/internal/application/services/tenant/remove-tenant"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantController struct {
	DbPool *pgxpool.Pool
}

func (c *TenantController) CreateTenant(writter http.ResponseWriter, request *http.Request) {
	var createTenantRequest createtenant.Request

	if err := http_router.ParseBodyToSchema(&createTenantRequest, request); err != nil {
		panic(err)
	}

	params := repositories.ParamsRs[createtenant.Request, *createtenant.Response, createtenant.CreateTenantService]{
		DbPool:  c.DbPool,
		New:     createtenant.New,
		Request: createTenantRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusCreated)
}

func (c *TenantController) RemoveTenant(writter http.ResponseWriter, request *http.Request) {
	var removeTenantRequest removetenant.Request

	if err := http_router.ParseBodyToSchema(&removeTenantRequest, request); err != nil {
		panic(err)
	}

	params := repositories.Params[removetenant.Request, removetenant.RemoveTenantService]{
		DbPool:  c.DbPool,
		New:     removetenant.New,
		Request: removeTenantRequest,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusCreated)
}
