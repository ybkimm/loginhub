package styles

import (
	"net/http"

	_ "embed"
)

//go:embed style.css
var styleCSS []byte

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(styleCSS)
}
