package session

import (
	"net/http"

	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Endpoint struct {
	DbPool *pgxpool.Pool
}

func (c *Endpoint) Http(writter http.ResponseWriter, request *http.Request) {
	authorizationHeader := request.Header.Get("Authorization")
	accessToken := authorizationHeader[len("Bearer "):]

	service := Handler{}

	response, err := service.Handler(request.Context(), Command{
		AccessToken: accessToken,
	})

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusOK)
}
