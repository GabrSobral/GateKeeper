package http_middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gate-keeper/internal/domain/errors"
	"github.com/go-chi/chi/v5/middleware"
)

type ErrorResponse struct {
	Title         string `json:"title"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlation_id"`
}

// WriteJSONError writes a JSON error response
func WriteJSONError(w http.ResponseWriter, statusCode int, title, message string, ctx context.Context) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	requestID := middleware.GetReqID(ctx)

	json.NewEncoder(w).Encode(ErrorResponse{Message: message, Title: title, CorrelationId: requestID})
}

// ErrorHandler is a middleware that recovers from panics and writes a JSON error response
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			ctx := r.Context()

			if err := recover(); err != nil {
				var statusCode int
				var message string
				var title string

				switch e := err.(type) {

				case *errors.InvalidRequestBodyResponse:
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)

					// requestID := middleware.GetReqID(ctx)

					json.NewEncoder(w).Encode(e)
					return

				case string:
					title = "Internal Server Error"
					message = e
					statusCode = http.StatusInternalServerError

				case error:
					err, ok := errors.ErrorsList[e.Error()]

					title = err.Title
					message = err.Message
					statusCode = err.Code

					if !ok {
						title = "Internal Server Error"
						message = e.Error()
						statusCode = http.StatusInternalServerError
					}
				default:
					title = "Internal Server Error"
					message = "Internal Server Error"
					statusCode = http.StatusInternalServerError
				}

				log.Printf("Recovered from panic: %v", err)
				WriteJSONError(w, statusCode, title, message, ctx)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
