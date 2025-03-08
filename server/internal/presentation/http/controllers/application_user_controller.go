package http_controllers

import (
	"fmt"
	"net/http"

	createapplicationuser "github.com/gate-keeper/internal/application/services/application-user/create-application-user"
	deleteapplicationuser "github.com/gate-keeper/internal/application/services/application-user/delete-application-user"
	getapplicationuserbyid "github.com/gate-keeper/internal/application/services/application-user/get-application-user-by-id"
	getuserbyemail "github.com/gate-keeper/internal/application/services/user/get-user-by-email"
	"github.com/gate-keeper/internal/infra/database/repositories"
	http_router "github.com/gate-keeper/internal/presentation/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicationUserController struct {
	DbPool *pgxpool.Pool
}

func (c *ApplicationUserController) CreateUser(writter http.ResponseWriter, request *http.Request) {
	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	var createApplicationUserRequest createapplicationuser.RequestBody

	if err := http_router.ParseBodyToSchema(&createApplicationUserRequest, request); err != nil {
		panic(err)
	}

	schema := createapplicationuser.Request{
		ApplicationID:         applicationIdUUID,
		DisplayName:           createApplicationUserRequest.DisplayName,
		FirstName:             createApplicationUserRequest.FirstName,
		LastName:              createApplicationUserRequest.LastName,
		Email:                 createApplicationUserRequest.Email,
		IsEmailConfirmed:      createApplicationUserRequest.IsEmailConfirmed,
		TemporaryPasswordHash: createApplicationUserRequest.TemporaryPasswordHash,
		IsMfaAuthAppEnabled:   createApplicationUserRequest.IsMfaAuthAppEnabled,
		IsMfaEmailEnabled:     createApplicationUserRequest.IsMfaEmailEnabled,
		Roles:                 createApplicationUserRequest.Roles,
	}

	params := repositories.ParamsRs[createapplicationuser.Request, *createapplicationuser.Response, createapplicationuser.CreateApplicationUserService]{
		DbPool:  c.DbPool,
		New:     createapplicationuser.New,
		Request: schema,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	http_router.SendJson(writter, response, http.StatusCreated)
}

func (c *ApplicationUserController) DeleteUser(writter http.ResponseWriter, request *http.Request) {
	applicationIDString := chi.URLParam(request, "applicationID")
	applicationIdUUID, err := uuid.Parse(applicationIDString)

	if err != nil {
		panic(err)
	}

	userIDString := chi.URLParam(request, "userID")
	userIdUUID, err := uuid.Parse(userIDString)

	if err != nil {
		panic(err)
	}

	schema := deleteapplicationuser.Request{
		ApplicationID: applicationIdUUID,
		UserID:        userIdUUID,
	}

	params := repositories.Params[deleteapplicationuser.Request, deleteapplicationuser.DeleteApplicationUserService]{
		DbPool:  c.DbPool,
		New:     deleteapplicationuser.New,
		Request: schema,
	}

	if err := repositories.WithTransaction(request.Context(), params); err != nil {
		panic(err)
	}

	http_router.SendJson(writter, nil, http.StatusNoContent)
}

func (c *ApplicationUserController) GetUserByEmailController(writter http.ResponseWriter, request *http.Request) {
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

func (c *ApplicationUserController) GetUserByIDController(writter http.ResponseWriter, request *http.Request) {
	userIDString := chi.URLParam(request, "userID")
	userIdUUID, err := uuid.Parse(userIDString)

	if err != nil {
		panic(err)
	}

	getUserByIDRequest := getapplicationuserbyid.Request{UserID: userIdUUID}

	params := repositories.ParamsRs[getapplicationuserbyid.Request, *getapplicationuserbyid.Response, getapplicationuserbyid.GetApplicationUserByID]{
		DbPool:  c.DbPool,
		New:     getapplicationuserbyid.New,
		Request: getUserByIDRequest,
	}

	response, err := repositories.WithTransactionRs(request.Context(), params)

	if err != nil {
		panic(err)
	}

	fmt.Println("asdasd askd alks jal jlaksj lkajs lkaj lkasd")
	fmt.Println(response)

	http_router.SendJson(writter, response, 200)
}
