package routes

import (
	"net/http"

	"github.com/ybkimm/loginhub"
	"github.com/ybkimm/loginhub/internal/assets"
	"go.uber.org/zap"
)

func imageHandler(logger *zap.Logger) http.Handler {
	return assets.NewHandlerWithStripPrefix(
		imagePrefix,
		assets.AssetHandlerOpts{
			FS:     loginhub.ImageDir(),
			Logger: logger,
		},
	)
}

func styleHandler(logger *zap.Logger) http.Handler {
	return assets.NewHandlerWithStripPrefix(
		stylePrefix,
		assets.AssetHandlerOpts{
			FS:     loginhub.StyleDir(),
			Logger: logger,
		},
	)
}

func scriptHandler(logger *zap.Logger) http.Handler {
	return assets.NewHandlerWithStripPrefix(
		scriptPrefix,
		assets.AssetHandlerOpts{
			FS:     loginhub.ScriptDir(),
			Logger: logger,
		},
	)
}
