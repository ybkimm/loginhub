package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/ybkimm/loginhub/internal/tpls"
)

const tpldir = "internal/tpls"
const cssfile = "internal/styles/style.css"

type realfs struct {
	base string
}

func (f *realfs) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(f.base, name))
}

func (f *realfs) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(f.base, name))
}

var mutex sync.Mutex
var tpl *tpls.Template
var css []byte

func reloadTemplate() error {
	mutex.Lock()
	defer mutex.Unlock()

	t, err := tpls.LoadFromFS(&realfs{tpldir})
	if err != nil {
		return err
	}

	tpl = t
	return nil
}

func watchFile(name string, cb func(event fsnotify.Event) error) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	watcher.Add(filepath.Join(wd, name))
	defer watcher.Close()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			log.Println(event)

			err = cb(event)
			if err != nil {
				return err
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}

			log.Fatalln(err)
		}
	}
}

func watchTpl() {
	watchFile(filepath.Join(tpldir, "html"), func(event fsnotify.Event) error {
		return reloadTemplate()
	})
}

func watchCSS() {
	watchFile(cssfile, func(event fsnotify.Event) error {
		data, err := os.ReadFile(cssfile)
		if err == os.ErrNotExist {
			css = nil
			return nil
		}
		if err != nil {
			return err
		}

		css = data

		return nil
	})
}

func init() {
	reloadTemplate()

	go watchTpl()
	go watchCSS()
}

func handle(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	p := r.URL.Path
	switch p {
	case "/assets/style.css":
		w.Header().Set("Content-Type", "text/css;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(css)

	case "/":
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Printf("Render Error: %v\n", err)
		}
	}
}

func main() {
	svr := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handle),
	}

	log.Fatal(svr.ListenAndServe())
}
