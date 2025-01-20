package http_middlewares

import (
	"context"
	"net/http"
	"strings"

	application_utils "github.com/gate-keeper/internal/application/utils"
)

func JwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			WriteJSONError(w, http.StatusUnauthorized, "Unauthorized", "Missing token", ctx)
			return
		}

		jwtToken := strings.Split(authHeader, "Bearer ")[1]
		isValid, userID, err := application_utils.ValidateToken(jwtToken)

		if !isValid {
			WriteJSONError(w, http.StatusUnauthorized, "Unauthorized", "Invalid token", ctx)
			return
		}

		if err != nil {
			WriteJSONError(w, http.StatusUnauthorized, "Token Validation Error", err.Error(), ctx)
			return
		}

		type contextKey string
		const userIDKey contextKey = "userId"

		// inject UserId on the request context
		ctx = context.WithValue(ctx, userIDKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
