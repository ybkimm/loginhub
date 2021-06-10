package main

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/ybkimm/loginhub/internal/tpls"
)

const tpldir = "internal/tpls"

var (
	tpl      *tpls.Template
	tplMutex sync.RWMutex
)

func init() {
	err := reloadTemplate()
	if err != nil {
		panic(err)
	}
	go watchTpl()
}

func reloadTemplate() error {
	tplMutex.Lock()
	defer tplMutex.Unlock()

	t, err := tpls.LoadFromFS(&realfs{tpldir})
	if err != nil {
		return err
	}

	tpl = t
	return nil
}

func watchTpl() {
	watchFile(filepath.Join(tpldir, "html"), func(event fsnotify.Event) error {
		return reloadTemplate()
	})
}

func handleTpl(w http.ResponseWriter, r *http.Request) {
	tplMutex.RLock()
	defer tplMutex.RUnlock()

	var (
		name = r.URL.Path[1:]
	)

	if len(name) == 0 {
		name = "index.html"
	}

	if !tpl.Has(name) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.Header().Set("Content-Length", "13")
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 Not Found")
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := tpl.ExecuteTemplate(w, name, nil)
	if err != nil {
		log.Printf("Render Error: %v\n", err)
	}
}
