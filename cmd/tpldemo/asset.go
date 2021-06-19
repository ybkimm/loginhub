package main

import (
	"net/http"

	"github.com/ybkimm/loginhub/internal/assets"
)

const (
	imagePrefix  = "/assets/images"
	imageDir     = "assets/images"
	stylePrefix  = "/assets/styles"
	styleDir     = "build/styles"
	scriptPrefix = "/assets/scripts"
	scriptDir    = "assets/scripts"
)

var (
	imageHandler = assets.NewHandlerWithStripPrefix(
		imagePrefix,
		assets.AssetHandlerOpts{
			FS:     &realfs{scriptDir},
			Logger: logger,
		},
	)
	styleHandler = assets.NewHandlerWithStripPrefix(
		stylePrefix,
		assets.AssetHandlerOpts{
			FS:       &realfs{styleDir},
			Logger:   logger,
			MimeType: "text/css;charset=utf-8",
		},
	)
	scriptHandler = http.StripPrefix(
		scriptPrefix,
		http.HandlerFunc(handleScript),
	)
)

func handleScript(w http.ResponseWriter, r *http.Request) {
	// TODO...
}
