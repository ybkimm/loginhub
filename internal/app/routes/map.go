package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func Handler() http.Handler {
	return router
}
