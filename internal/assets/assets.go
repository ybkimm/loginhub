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
	fs     fs.FS
	logger *zap.Logger
}

func NewHandler(fs fs.FS, logger *zap.Logger) *AssetHandler {
	return &AssetHandler{
		fs:     fs,
		logger: logger,
	}
}

func (h *AssetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		filename = r.URL.Path[1:]
		filetype = mime.TypeByExtension(filepath.Ext(filename))
	)

	if len(filetype) == 0 {
		filetype = "application/octet-stream"
	}

	h.logger.Debug("assets: trying to open", zap.String("filename", filename), zap.String("mime type", filetype))

	f, err := h.fs.Open(filename)
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

	w.Header().Set("Content-Type", filetype)
	w.Header().Set("Content-Length", fileSize)
	w.WriteHeader(http.StatusOK)
	io.Copy(w, f)
}

func sendInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "500 Internal Server Error")
}
