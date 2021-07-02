package routes

import (
	"net/http"
	"strconv"

	"github.com/ybkimm/loginhub"
	"go.uber.org/zap"
)

var viewContentLength string

func init() {
	viewContentLength = strconv.FormatInt(int64(len(loginhub.IndexHTML)), 10)
}

// viewHandler does simple thing - write index.html to response.
// This method can be used for every view!
func viewHandler(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		w.Header().Set("Content-Length", viewContentLength)
		w.WriteHeader(http.StatusOK)
		w.Write(loginhub.IndexHTML)
	}
}
