package routes

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	assetPrefix  = "/assets"
	imagePrefix  = assetPrefix + "/images"
	stylePrefix  = assetPrefix + "/styles"
	scriptPrefix = assetPrefix + "/scripts"
)

func NewRouter(logger *zap.Logger) *mux.Router {
	r := mux.NewRouter()

	r.Methods("GET").
		Path("/health").
		Handler(healthHandler(logger))

	r.Methods("GET").
		PathPrefix(imagePrefix).
		Handler(imageHandler(logger))

	r.Methods("GET").
		PathPrefix(stylePrefix).
		Handler(styleHandler(logger))

	r.Methods("GET").
		PathPrefix(scriptPrefix).
		Handler(scriptHandler(logger))

	r.NotFoundHandler = viewHandler(logger)

	return r
}
