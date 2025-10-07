package app

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestApp_Run(t *testing.T) {
	tests := []struct {
		name         string
		config       string
		port         int
		expectError  bool
		expectedAddr string
	}{
		{
			name: "valid_config_default_port",
			config: `{
				"endpoints": [
					{
						"url": "/api/users",
						"response": {
							"id": "uuid",
							"name": "name",
							"email": "email"
						}
					}
				],
				"cache": 5
			}`,
			port:         0,
			expectError:  false,
			expectedAddr: ":0",
		},
		{
			name: "valid_config_custom_port",
			config: `{
				"endpoints": [
					{
						"url": "/api/products",
						"response": {
							"id": "uuid",
							"title": "word",
							"price": "float"
						}
					}
				],
				"cache": 10
			}`,
			port:         8080,
			expectError:  false,
			expectedAddr: ":8080",
		},
		{
			name: "invalid_json_syntax",
			config: `{
				"endpoints": [
					{
						"url": "/api/test",
						"response": {
							"id": "uuid"
						}
					}
				],
				"cache": 1
			`,
			port:        3000,
			expectError: true,
		},
		{
			name: "empty_endpoints",
			config: `{
				"endpoints": [],
				"cache": 1
			}`,
			port:         4000,
			expectError:  false,
			expectedAddr: ":4000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfgPath := filepath.Join(t.TempDir(), "config.json")
			if err := os.WriteFile(cfgPath, []byte(tt.config), 0o600); err != nil {
				t.Fatalf("Failed to write config: %v", err)
			}

			srv, err := Run(cfgPath, tt.port)

			if tt.expectError {
				if err == nil {
					t.Fatal("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if srv.Addr != tt.expectedAddr {
				t.Errorf("Expected address %s, got %s", tt.expectedAddr, srv.Addr)
			}

			if srv.Handler == nil {
				t.Error("Expected handler to be set")
			}
		})
	}
}

func TestApp_RunMissingConfig(t *testing.T) {
	missingPath := filepath.Join(t.TempDir(), "nonexistent.json")
	_, err := Run(missingPath, 8080)
	if err == nil {
		t.Fatal("Expected error for missing config file")
	}
}

func TestApp_RunLifecycle(t *testing.T) {
	cfgPath := filepath.Join(t.TempDir(), "config.json")
	config := `{
		"endpoints": [
			{
				"url": "/api/health",
				"response": {
					"status": "word",
					"timestamp": "date"
				}
			}
		],
		"cache": 3
	}`

	if err := os.WriteFile(cfgPath, []byte(config), 0o600); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	srv, err := Run(cfgPath, 0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Server error: %v", err)
		}
	}()

	time.Sleep(50 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		t.Errorf("Failed to shutdown server: %v", err)
	}
}
