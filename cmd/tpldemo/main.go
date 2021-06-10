package main

import (
	"log"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[HTTP] %s %s", r.Method, r.URL.Path)

	switch {
	case strings.HasPrefix(r.URL.Path, imagePrefix):
		imageHandler.ServeHTTP(w, r)

	case strings.HasPrefix(r.URL.Path, scriptPrefix):
		scriptHandler.ServeHTTP(w, r)

	case strings.HasPrefix(r.URL.Path, stylePrefix):
		styleHandler.ServeHTTP(w, r)

	default:
		handleTpl(w, r)
	}
}

func main() {
	svr := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handle),
	}

	log.Fatal(svr.ListenAndServe())
}
