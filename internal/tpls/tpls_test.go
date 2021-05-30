package tpls

import (
	"bytes"
	"testing"
	"testing/fstest"
)

func TestFS(t *testing.T) {
	err := fstest.TestFS(tplfs, "html/_template.html", "html/index.html")
	if err != nil {
		t.Errorf("looks like the required files in the template directory are missing: %s", err)
	}
}

var fakeTpls = fstest.MapFS{
	"html/_template.html": &fstest.MapFile{
		Data: []byte(`{{define "template"}}<template>{{block "contents" .}}No Contents!{{end}}</template>{{end}}`),
	},
	"html/_nav.html": &fstest.MapFile{
		Data: []byte(`{{define "nav"}}<nav>NAV</nav>{{end}}`),
	},
	"html/index.html": &fstest.MapFile{
		Data: []byte(`{{define "contents"}}{{template "nav" .}}Hello, World!{{end}}{{template "template" .}}`),
	},
}

const expectedTemplateResult = "<template><nav>NAV</nav>Hello, World!</template>"

func TestReadTpls(t *testing.T) {
	tpl, err := readTpls(fakeTpls)
	if err != nil {
		t.Errorf("readTpls returned an error: %s", err)
		return
	}

	buf := new(bytes.Buffer)

	err = tpl.ExecuteTemplate(buf, "???", nil)
	if err != ErrTemplateNotFound {
		t.Errorf("ExecuteTemplate throws an error: %s, want: %s", err, ErrTemplateNotFound)
	}

	err = tpl.ExecuteTemplate(buf, "index.html", nil)
	if err != nil {
		t.Errorf("ExecuteTemplate returned an error: %s", err)
		return
	}

	if !bytes.Equal(buf.Bytes(), []byte(expectedTemplateResult)) {
		t.Errorf("ExecuteTemplate returned wrong result: %s, want: %s", buf.String(), expectedTemplateResult)
		return
	}
}
