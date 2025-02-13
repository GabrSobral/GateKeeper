package http_controllers

import (
	"net/http"

	createrole "github.com/gate-keeper/internal/application/services/application-role/create-role"
	deleterole "github.com/gate-keeper/internal/application/services/application-role/delete-role"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationRoleController struct {
	DbPool *pgxpool.Pool
}

func (c *ApplicationRoleController) CreateRole(writter http.ResponseWriter, request *http.Request) {
	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	var schema createrole.Request

	if err := http_router.ParseBodyToSchema(&schema, request); err != nil {
		panic(err)
	}

	schema.ApplicationID = applicationIdUUID

	params := repositories.ParamsRs[createrole.Request, *createrole.Response, createrole.CreateRoleService]{
		DbPool:  c.DbPool,
		New:     createrole.New,
		Request: schema,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusCreated)
}

func (c *ApplicationRoleController) RemoveRole(writter http.ResponseWriter, request *http.Request) {
	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	roleIDString := chi.URLParam(request, "roleID")
	roleIdUUID, err := uuid.Parse(roleIDString)

	if err != nil {
		panic(err)
	}

	var schema deleterole.Request

	if err := http_router.ParseBodyToSchema(&schema, request); err != nil {
		panic(err)
	}

	schema.ApplicationID = applicationIdUUID
	schema.RoleID = roleIdUUID

	params := repositories.Params[deleterole.Request, deleterole.DeleteRoleService]{
		DbPool:  c.DbPool,
		New:     deleterole.New,
		Request: schema,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}
