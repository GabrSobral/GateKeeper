package http_controllers

import (
	"log"
	"net/http"

	getuserbyemail "github.com/gate-keeper/internal/application/services/user/get-user-by-email"
	getuserbyid "github.com/gate-keeper/internal/application/services/user/get-user-by-id"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
	DbPool *pgxpool.Pool
}

func (c *UserController) GetUserByEmailController(writter http.ResponseWriter, request *http.Request) {
	userEmailString := chi.URLParam(request, "email")
	getUserByEmailRequest := getuserbyemail.Request{Email: userEmailString}

	params := repositories.ParamsRs[getuserbyemail.Request, *getuserbyemail.Response, getuserbyemail.GetUserByEmail]{
		DbPool:  c.DbPool,
		New:     getuserbyemail.New,
		Request: getUserByEmailRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, 200)
}

func (c *UserController) GetUserByIDController(writter http.ResponseWriter, request *http.Request) {
	log.Printf("asd;askdl;as")

	userIDString := chi.URLParam(request, "userID")

	log.Printf("UUID here: %v", userIDString)

	userIdUUID, err := uuid.Parse(userIDString)

	if err != nil {
		panic(err)
	}

	getUserByIDRequest := getuserbyid.Request{UserID: userIdUUID}

	params := repositories.ParamsRs[getuserbyid.Request, *getuserbyid.Response, getuserbyid.GetUserByID]{
		DbPool:  c.DbPool,
		New:     getuserbyid.New,
		Request: getUserByIDRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, 200)
}
