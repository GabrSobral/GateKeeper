package http_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/go-playground/validator/v10"
)

func SendJson(w http.ResponseWriter, data any, status int) {
	jsonData, err := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error on trying to parse JSON", http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Error on trying to parse JSON" }`))
	}

	w.WriteHeader(status)
	w.Write(jsonData)
}

func ParseBodyToSchema[T any](schema *T, request *http.Request) error {
	if err := json.NewDecoder(request.Body).Decode(&schema); err != nil {
		if err := validateStruct[T](schema); err != nil {
			return err
		}

		return err
	}

	if err := validateStruct[T](schema); err != nil {
		panic(err)
	}

	return nil
}

func validateStruct[T any](instance *T) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(instance)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		invalidBodyFields := []errors.InvalidRequestBodyField{}

		for _, err := range err.(validator.ValidationErrors) {
			invalidBodyFields = append(invalidBodyFields, errors.InvalidRequestBodyField{
				Field:   err.Field(),
				Type:    err.Type().Name(),
				Message: err.Error(),
			})
		}

		return &errors.InvalidRequestBodyResponse{
			Title:   "Invalid Request Body",
			Message: "Invalid Request Body, check the fields below",
			Fields:  invalidBodyFields,
		}
	}

	return nil
}
