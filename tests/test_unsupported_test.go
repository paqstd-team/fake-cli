package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
	"github.com/paqstd-team/fake-cli/handler"
)

func TestData_UnsupportedTypePath(t *testing.T) {
	cfg := config.Config{Endpoints: []config.Endpoint{{URL: "/u", Fields: 123, Response: "single"}}}
	h := handler.MakeHandler(cfg)
	req := httptest.NewRequest(http.MethodGet, "/u", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	var s string
	if err := json.Unmarshal(w.Body.Bytes(), &s); err != nil {
		t.Fatalf("json: %v", err)
	}
	if s == "" {
		t.Fatalf("expected unsupported type message, got empty")
	}
}
