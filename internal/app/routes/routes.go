package routes

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger) *mux.Router {
	r := mux.NewRouter()

	r.Methods("GET").
		Path("/health").
		Handler(healthHandler(logger))

	r.Methods("GET").
		Path("/login").
		Handler(viewHandler(logger))

	return r
}
