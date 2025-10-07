package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
	"github.com/paqstd-team/fake-cli/handler"
)

func TestHandler_SingleAndList(t *testing.T) {
	cfg := config.Config{Endpoints: []config.Endpoint{
		{URL: "/single", Fields: map[string]any{"id": "uuid", "name": "name"}, Response: "single"},
		{URL: "/list", Fields: map[string]any{"id": "uuid"}, Response: "list"},
	}, Cache: 2}
	h := handler.MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/single", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	var single map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &single); err != nil {
		t.Fatalf("json: %v", err)
	}
	if _, ok := single["id"]; !ok {
		t.Fatalf("missing id")
	}

	u := &url.URL{Path: "/list"}
	q := u.Query()
	q.Set("page", "1")
	q.Set("per_page", "2")
	u.RawQuery = q.Encode()
	req = httptest.NewRequest(http.MethodGet, u.String(), nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	var list []any
	if err := json.Unmarshal(w.Body.Bytes(), &list); err != nil {
		t.Fatalf("json list: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("len: %d", len(list))
	}

	// cache-hit path
	req = httptest.NewRequest(http.MethodGet, "/single", nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	first := w.Body.String()
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	second := w.Body.String()
	if first != second {
		t.Fatalf("expected cache hit to return same body")
	}
}
