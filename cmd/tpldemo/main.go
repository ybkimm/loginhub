package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	switch {
	case strings.HasPrefix(p, "/public/"):
		assetHandler.ServeHTTP(w, r)

	case p == "/":
		handleTpl(w, r)

	default:
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.Header().Set("Content-Length", "13")
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 Not Found")
	}
}

func main() {
	svr := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handle),
	}

	log.Fatal(svr.ListenAndServe())
}
