package getapplicationbyid

import (
	"net/http"

	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Endpoint struct {
	DbPool *pgxpool.Pool
}

func (c *Endpoint) Http(writter http.ResponseWriter, request *http.Request) {
	organizationIDString := chi.URLParam(request, "organizationID")
	organizationIdUUID, err := uuid.Parse(organizationIDString)

	if err != nil {
		panic(err)
	}

	applicationIDString := chi.URLParam(request, "applicationID")
	applicationnIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	requestSchema := Query{
		OrganizationID: organizationIdUUID,
		ApplicationID:  applicationnIdUUID,
	}

	params := repositories.ParamsRs[Query, *Response, Handler]{
		DbPool:  c.DbPool,
		New:     New,
		Request: requestSchema,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}
