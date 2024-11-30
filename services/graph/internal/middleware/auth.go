package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

type contextKey string

const (
	userIDKey contextKey = "userID"
	userID    string     = "2e597e90-ece5-463e-8608-ff687bf286da" // TODO: authentication
)

// Route middleware for the entire request (NOT graphql interceptor middleware)
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("user", "id", userID)

		ctx := WithUserID(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserID returns the user ID from the context.
// Panics if the user ID is not set.
func GetUserID(ctx context.Context) string {
	user, ok := ctx.Value(userIDKey).(string)
	if !ok {
		panic("user ID not set")
	}
	return user
}
