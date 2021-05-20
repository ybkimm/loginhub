package styles

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStyleHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/assets/style.css", nil)
	rec := httptest.NewRecorder()

	StyleHandler(rec, req)

	if contentType := rec.Header().Get("Content-Type"); contentType != "text/css;charset=utf-8" {
		t.Errorf("wrong content-type: %s, want text/css;charset=utf-8", contentType)
		return
	}

	if rec.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: %d, want 200", rec.Code)
		return
	}

	if body := rec.Body.Bytes(); !bytes.Equal(body, styleCSS) {
		t.Errorf("handler returned wrong body: %s, want %s", body, styleCSS)
	}
}
