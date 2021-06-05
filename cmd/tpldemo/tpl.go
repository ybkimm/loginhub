package main

import (
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

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("Render Error: %v\n", err)
	}
}
