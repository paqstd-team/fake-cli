package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfig_Load(t *testing.T) {
	tests := []struct {
		name              string
		config            string
		expectError       bool
		expectedEndpoints int
		expectedCache     int
	}{
		{
			name: "valid_api_config",
			config: `{
				"endpoints": [
					{
						"url": "/api/users",
						"type": "GET",
						"response": {
							"id": "uuid",
							"name": "name",
							"email": "email",
							"created_at": "date"
						},
						"status": 200,
						"cache": 10
					},
					{
						"url": "/api/users",
						"type": "POST",
						"payload": {
							"name": "",
							"email": ""
						},
						"response": {
							"id": "uuid",
							"name": "name",
							"email": "email"
						},
						"status": 201
					}
				]
			}`,
			expectError:       false,
			expectedEndpoints: 2,
			expectedCache:     0,
		},
		{
			name: "minimal_config",
			config: `{
				"endpoints": [
					{
						"url": "/health",
						"response": {
							"status": "word"
						}
					}
				]
			}`,
			expectError:       false,
			expectedEndpoints: 1,
			expectedCache:     0,
		},
		{
			name: "empty_endpoints",
			config: `{
				"endpoints": []
			}`,
			expectError:       false,
			expectedEndpoints: 0,
			expectedCache:     0,
		},
		{
			name: "complex_response_structure",
			config: `{
				"endpoints": [
					{
						"url": "/api/products",
						"response": [
							{
								"id": "uuid",
								"title": "word",
								"price": "float",
								"category": {
									"id": "int",
									"name": "word"
								},
								"tags": ["word", "word", "word"]
							}
						],
						"cache": 20
					}
				]
			}`,
			expectError:       false,
			expectedEndpoints: 1,
			expectedCache:     0,
		},
		{
			name: "malformed_json",
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
			expectError: true,
		},
		{
			name: "invalid_json_syntax",
			config: `{
				"endpoints": [
					{
						"url": "/api/test",
						"response": {
							"id": "uuid",
						}
					}
				],
				"cache": 1
			}`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			configPath := filepath.Join(dir, "config.json")

			if err := os.WriteFile(configPath, []byte(tt.config), 0o600); err != nil {
				t.Fatalf("Failed to write config file: %v", err)
			}

			cfg, err := LoadConfigFromFile(configPath)

			if tt.expectError {
				if err == nil {
					t.Fatal("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(cfg.Endpoints) != tt.expectedEndpoints {
				t.Errorf("Expected %d endpoints, got %d", tt.expectedEndpoints, len(cfg.Endpoints))
			}

			// Cache is now per-endpoint, so we don't check global cache
		})
	}
}

func TestConfig_FileNotFound(t *testing.T) {
	_, err := LoadConfigFromFile("/nonexistent/path/config.json")
	if err == nil {
		t.Fatal("Expected error for nonexistent file")
	}
}

func TestConfig_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "empty.json")

	if err := os.WriteFile(configPath, []byte(""), 0o600); err != nil {
		t.Fatalf("Failed to write empty file: %v", err)
	}

	_, err := LoadConfigFromFile(configPath)
	if err == nil {
		t.Fatal("Expected error for empty file")
	}
}

func TestConfig_InvalidFilePermissions(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "restricted.json")

	config := `{
		"endpoints": [
			{
				"url": "/api/test",
				"response": {
					"id": "uuid"
				}
			}
		],
		"cache": 1
	}`

	if err := os.WriteFile(configPath, []byte(config), 0o000); err != nil {
		t.Fatalf("Failed to write restricted file: %v", err)
	}

	_, err := LoadConfigFromFile(configPath)
	if err == nil {
		t.Fatal("Expected error for file with no read permissions")
	}
}
