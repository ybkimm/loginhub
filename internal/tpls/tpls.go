package tpls

import (
	"embed"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
)

//go:embed html/*.html
var tplfs embed.FS

var ErrTemplateNotFound = errors.New("tpls: template not found")

type Template struct {
	templates map[string]*template.Template
}

func Load() (*Template, error) {
	return readTpls(tplfs)
}

func LoadFromFS(tplfs fs.ReadDirFS) (*Template, error) {
	return readTpls(tplfs)
}

func readTpls(tplfs fs.ReadDirFS) (*Template, error) {
	baseTpl := template.New("")

	baseTpl.Funcs(funcs)

	// Get the files that need to be read from the directory
	files, err := tplfs.ReadDir("html")
	if err != nil {
		return nil, err
	}

	// Add files starting with "_" to the base template and
	// exclude them from the file list.
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		fileName := file.Name()
		if fileName[0] != '_' {
			continue
		}

		_, err = parseTpl(baseTpl, tplfs, filepath.Join("html/", fileName))
		if err != nil {
			return nil, err
		}

		files[i] = files[len(files)-1]
		files = files[:len(files)-1]
	}

	// Template map is here...
	var tpls = make(map[string]*template.Template)

	// Parse the files and add them to the template map.
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		var fileName = file.Name()

		tpl, err := baseTpl.Clone()
		if err != nil {
			return nil, err
		}
		tpl = tpl.New(fileName)

		_, err = parseTpl(tpl, tplfs, filepath.Join("html/", fileName))
		if err != nil {
			return nil, err
		}

		tpls[fileName] = tpl
	}

	return &Template{templates: tpls}, nil
}

func parseTpl(t *template.Template, fs fs.FS, name string) (*template.Template, error) {
	file, err := fs.Open(name)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return t.Parse(string(data))
}

func (t *Template) ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	tpl, ok := t.templates[name]
	if !ok {
		return ErrTemplateNotFound
	}
	return tpl.Execute(w, data)
}
