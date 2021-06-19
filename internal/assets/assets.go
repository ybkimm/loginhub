package assets

import (
	"errors"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"go.uber.org/zap"
)

type AssetHandler struct {
	fs       fs.FS
	logger   *zap.Logger
	mimeType string // Infer from extension if empty string
}

type AssetHandlerOpts struct {
	FS       fs.FS
	Logger   *zap.Logger
	MimeType string
}

func NewHandler(opts AssetHandlerOpts) http.Handler {
	return &AssetHandler{
		fs:       opts.FS,
		logger:   opts.Logger,
		mimeType: opts.MimeType,
	}
}

func NewHandlerWithStripPrefix(prefix string, opts AssetHandlerOpts) http.Handler {
	return http.StripPrefix(
		prefix,
		NewHandler(opts),
	)
}

func (h *AssetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		fileName = r.URL.Path[1:]
		fileType string
	)

	if len(h.mimeType) == 0 {
		fileType = mime.TypeByExtension(filepath.Ext(fileName))
		if len(fileType) == 0 {
			fileType = "application/octet-stream"
		}
	} else {
		fileType = h.mimeType
	}

	h.logger.Debug("assets: trying to open", zap.String("filename", fileName), zap.String("mime type", fileType))

	f, err := h.fs.Open(fileName)
	if errors.Is(err, fs.ErrNotExist) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 Not Found")
		return
	} else if err != nil {
		h.logger.Error("assets: open error", zap.Error(err))
		sendInternalServerError(w)
		return
	}

	stat, err := f.Stat()
	if err != nil {
		h.logger.Error("assets: stat error", zap.Error(err))
		sendInternalServerError(w)
		return
	}

	var fileSize = strconv.FormatInt(stat.Size(), 10)

	w.Header().Set("Content-Type", fileType)
	w.Header().Set("Content-Length", fileSize)
	w.WriteHeader(http.StatusOK)
	io.Copy(w, f)
}

func sendInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "500 Internal Server Error")
}
