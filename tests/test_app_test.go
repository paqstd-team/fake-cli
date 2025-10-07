package tests

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paqstd-team/fake-cli/app"
)

func TestApp_RunLifecycle(t *testing.T) {
	cfgPath := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(cfgPath, []byte(`{"endpoints":[{"url":"/ping","fields":{"id":"uuid"},"response":"single"}],"cache":1}`), 0o600); err != nil {
		t.Fatalf("write cfg: %v", err)
	}
	srv, err := app.Run(cfgPath, 0)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	go func() { _ = srv.ListenAndServe() }()
	time.Sleep(30 * time.Millisecond)
	_ = srv.Close()
	_ = srv.Shutdown(context.Background())
	_ = http.ErrServerClosed
}

func TestApp_RunMissingConfig(t *testing.T) {
	_, err := app.Run(filepath.Join(t.TempDir(), "missing.json"), 0)
	if err == nil {
		t.Fatalf("expected error for missing config file")
	}
}
