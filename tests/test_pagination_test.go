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

func TestPagination_DefaultsAndParsing(t *testing.T) {
	cfg := config.Config{Endpoints: []config.Endpoint{{URL: "/list", Fields: map[string]any{"id": "uuid"}, Response: "list"}}}
	h := handler.MakeHandler(cfg)

	// defaults
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	var items []any
	if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(items) != 10 {
		t.Fatalf("default per_page len: %d", len(items))
	}

	// custom
	u := &url.URL{Path: "/list"}
	q := u.Query()
	q.Set("page", "3")
	q.Set("per_page", "7")
	u.RawQuery = q.Encode()
	req = httptest.NewRequest(http.MethodGet, u.String(), nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(items) != 7 {
		t.Fatalf("custom per_page len: %d", len(items))
	}

	// page=0 -> empty
	u = &url.URL{Path: "/list"}
	q = u.Query()
	q.Set("page", "0")
	u.RawQuery = q.Encode()
	req = httptest.NewRequest(http.MethodGet, u.String(), nil)
	w = httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
		t.Fatalf("json: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected empty list when page=0, got %d", len(items))
	}
}
