package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
	"github.com/paqstd-team/fake-cli/handler"
)

// Force json.Marshal error with non-finite float values are allowed, so craft a channel which is not marshalable via generateData path
func TestHandler_JSONMarshalError(t *testing.T) {
	// inject marshal failure
	original := handler.JSONMarshal
	handler.JSONMarshal = func(v any) ([]byte, error) { return nil, errors.New("boom") }
	defer func() { handler.JSONMarshal = original }()

	cfg := config.Config{Endpoints: []config.Endpoint{{URL: "/err", Fields: map[string]any{"ok": "word"}, Response: "single"}}}
	h := handler.MakeHandler(cfg)
	req := httptest.NewRequest(http.MethodGet, "/err", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
