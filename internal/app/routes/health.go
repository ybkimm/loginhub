package routes

import (
	"net/http"

	"go.uber.org/zap"
)

func healthHandler(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", "15")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}
}
