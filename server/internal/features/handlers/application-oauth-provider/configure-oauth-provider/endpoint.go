package configureoauthprovider

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
	var requestBody RequestBody

	if err := http_router.ParseBodyToSchema(&requestBody, request); err != nil {
		panic(err)
	}

	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	command := Command{
		ApplicationID: applicationIdUUID,
		Name:          requestBody.Name,
		ClientID:      requestBody.ClientID,
		ClientSecret:  requestBody.ClientSecret,
		RedirectURI:   requestBody.RedirectURI,
		Enabled:       requestBody.Enabled,
	}

	params := repositories.Params[Command, Handler]{
		DbPool:  c.DbPool,
		New:     New,
		Request: command,
	}

	errHandler := repositories.WithTransaction(request.Context(), params)

	if errHandler != nil {
		panic(errHandler)
	}

	http_router.SendJson(writter, nil, http.StatusCreated)
}
