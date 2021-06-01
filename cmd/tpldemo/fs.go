package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type realfs struct {
	base string
}

var _ fs.FS = (*realfs)(nil)

func (f *realfs) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(f.base, name))
}

func (f *realfs) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(f.base, name))
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
