package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	getuserbyemail "github.com/guard-service/internal/application/services/user/get-user-by-email"
	getuserbyid "github.com/guard-service/internal/application/services/user/get-user-by-id"
	utils "github.com/guard-service/internal/presentation/http"
)

type UserController struct {
	GetUserByEmailService *getuserbyemail.GetUserByEmail
	GetUserByIDService    *getuserbyid.GetUserByID
}

func (c *UserController) GetUserByEmailController(writter http.ResponseWriter, request *http.Request) {
	userEmailString := chi.URLParam(request, "email")
	getUserByEmailRequest := getuserbyemail.Request{Email: userEmailString}

	response, err := c.GetUserByEmailService.Handler(request.Context(), getUserByEmailRequest)

	if err != nil {
		panic(err)
	}

	utils.SendJson(writter, response, 200)
}

func (c *UserController) GetUserByIDController(writter http.ResponseWriter, request *http.Request) {
	userIDString := chi.URLParam(request, "userID")

	userIdUUID, err := uuid.Parse(userIDString)

	if err != nil {
		panic(err)
	}

	getUserByIDRequest := getuserbyid.Request{UserID: userIdUUID}

	response, err := c.GetUserByIDService.Handler(request.Context(), getUserByIDRequest)

	if err != nil {
		panic(err)
	}

	utils.SendJson(writter, response, 200)
}
