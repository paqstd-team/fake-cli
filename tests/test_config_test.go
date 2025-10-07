package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
)

func TestConfig_Load(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "cfg.json")
	content := `{"endpoints":[{"url":"/u","fields":{"id":"uuid"},"response":"single"}],"cache":5}`
	if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	cfg, err := config.LoadConfigFromFile(p)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if len(cfg.Endpoints) != 1 || cfg.Cache != 5 {
		t.Fatalf("cfg: %+v", cfg)
	}
}

func TestConfig_LoadMalformed(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "bad.json")
	content := `{"endpoints": [ { "url": "/u", } ]}` // malformed JSON
	if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err := config.LoadConfigFromFile(p); err == nil {
		t.Fatalf("expected error for malformed json")
	}
}
