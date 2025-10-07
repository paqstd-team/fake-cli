package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/paqstd-team/fake-cli/v2/config"
)

func TestHandler_HTTPMethods(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       config.Endpoint
		method         string
		expectedStatus int
		expectError    bool
	}{
		{
			name: "GET_endpoint_success",
			endpoint: config.Endpoint{
				URL:      "/api/users",
				Type:     http.MethodGet,
				Response: map[string]any{"id": "uuid", "name": "name"},
				Status:   200,
			},
			method:         http.MethodGet,
			expectedStatus: 200,
			expectError:    false,
		},
		{
			name: "POST_endpoint_success",
			endpoint: config.Endpoint{
				URL:      "/api/users",
				Type:     http.MethodPost,
				Response: map[string]any{"id": "uuid", "status": "created"},
				Status:   201,
			},
			method:         http.MethodPost,
			expectedStatus: 201,
			expectError:    false,
		},
		{
			name: "DELETE_endpoint_success",
			endpoint: config.Endpoint{
				URL:      "/api/users/123",
				Type:     http.MethodDelete,
				Response: nil,
				Status:   204,
			},
			method:         http.MethodDelete,
			expectedStatus: 204,
			expectError:    false,
		},
		{
			name: "PUT_endpoint_success",
			endpoint: config.Endpoint{
				URL:      "/api/users/123",
				Type:     http.MethodPut,
				Response: map[string]any{"id": "uuid", "name": "name", "updated": true},
				Status:   200,
			},
			method:         http.MethodPut,
			expectedStatus: 200,
			expectError:    false,
		},
		{
			name: "method_not_allowed",
			endpoint: config.Endpoint{
				URL:      "/api/users",
				Type:     http.MethodGet,
				Response: map[string]any{"id": "uuid"},
			},
			method:         http.MethodPost,
			expectedStatus: 405,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{tt.endpoint},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(tt.method, tt.endpoint.URL, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.endpoint.Response != nil && tt.expectedStatus < 400 {
				var response map[string]any
				if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
			}
		})
	}
}

func TestHandler_ResponseTypes(t *testing.T) {
	tests := []struct {
		name         string
		endpoint     config.Endpoint
		expectedType string
	}{
		{
			name: "single_object_response",
			endpoint: config.Endpoint{
				URL: "/api/user",
				Response: map[string]any{
					"id":    "uuid",
					"name":  "name",
					"email": "email",
				},
			},
			expectedType: "object",
		},
		{
			name: "array_response",
			endpoint: config.Endpoint{
				URL: "/api/users",
				Response: []any{
					map[string]any{"id": "uuid", "name": "name"},
					map[string]any{"id": "uuid", "name": "name"},
				},
			},
			expectedType: "array",
		},
		{
			name: "empty_array_response",
			endpoint: config.Endpoint{
				URL:      "/api/empty",
				Response: []any{},
			},
			expectedType: "array",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{tt.endpoint},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(http.MethodGet, tt.endpoint.URL, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var response interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			switch tt.expectedType {
			case "object":
				if _, ok := response.(map[string]any); !ok {
					t.Errorf("Expected object response, got %T", response)
				}
			case "array":
				if _, ok := response.([]any); !ok {
					t.Errorf("Expected array response, got %T", response)
				}
			}
		})
	}
}

func TestHandler_PayloadValidation(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       config.Endpoint
		payload        string
		expectedStatus int
	}{
		{
			name: "valid_payload",
			endpoint: config.Endpoint{
				URL:      "/api/users",
				Type:     http.MethodPost,
				Payload:  map[string]any{"name": "", "email": ""},
				Response: map[string]any{"id": "uuid", "name": "name", "email": "email"},
				Status:   201,
			},
			payload:        `{"name":"John Doe","email":"john@example.com"}`,
			expectedStatus: 201,
		},
		{
			name: "missing_required_field",
			endpoint: config.Endpoint{
				URL:      "/api/users",
				Type:     http.MethodPost,
				Payload:  map[string]any{"name": "", "email": ""},
				Response: map[string]any{"id": "uuid"},
				Status:   400,
			},
			payload:        `{"name":"John Doe"}`,
			expectedStatus: 400,
		},
		{
			name: "invalid_json",
			endpoint: config.Endpoint{
				URL:      "/api/users",
				Type:     http.MethodPost,
				Payload:  map[string]any{"name": ""},
				Response: map[string]any{"id": "uuid"},
				Status:   400,
			},
			payload:        `{"name":"John Doe"`,
			expectedStatus: 400,
		},
		{
			name: "non_map_body",
			endpoint: config.Endpoint{
				URL:      "/api/users",
				Type:     http.MethodPost,
				Payload:  map[string]any{"name": ""},
				Response: map[string]any{"id": "uuid"},
				Status:   400,
			},
			payload:        `"not a map"`,
			expectedStatus: 400,
		},
		{
			name: "array_payload_valid",
			endpoint: config.Endpoint{
				URL:      "/api/bulk",
				Type:     http.MethodPost,
				Payload:  []any{map[string]any{"id": "string"}},
				Response: map[string]any{"processed": "int"},
				Status:   200,
			},
			payload:        `[{"id":"123"},{"id":"456"}]`,
			expectedStatus: 200,
		},
		{
			name: "array_payload_invalid",
			endpoint: config.Endpoint{
				URL:      "/api/bulk",
				Type:     http.MethodPost,
				Payload:  []any{map[string]any{"id": "string"}},
				Response: map[string]any{"processed": "int"},
				Status:   400,
			},
			payload:        `[{"id":"123"},{"invalid":"456"}]`,
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{tt.endpoint},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(tt.endpoint.Type, tt.endpoint.URL, strings.NewReader(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_CacheBehavior(t *testing.T) {
	cacheSize := 2
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL: "/api/cached",
				Response: map[string]any{
					"id":        "uuid",
					"timestamp": "date",
				},
				Cache: &cacheSize,
			},
		},
	}
	handler := MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/cached", nil)
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req)

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req)

	if w1.Body.String() != w2.Body.String() {
		t.Error("Expected cached responses to be identical")
	}
}

func TestHandler_JSONMarshalError(t *testing.T) {
	original := JSONMarshal
	JSONMarshal = func(v any) ([]byte, error) { return nil, errors.New("marshal error") }
	defer func() { JSONMarshal = original }()

	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL:      "/api/error",
				Response: map[string]any{"data": "word"},
			},
		},
	}
	handler := MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/error", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestHandler_ReadBodyError(t *testing.T) {
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL:      "/api/test",
				Type:     http.MethodPost,
				Payload:  map[string]any{"field": "string"},
				Response: map[string]any{"ok": "word"},
			},
		},
	}
	handler := MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodPost, "/api/test", &errorReader{})
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestHandler_HTTPMethodsAndStatusCodes(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       config.Endpoint
		method         string
		expectedStatus int
		description    string
	}{
		{
			name: "PUT_method",
			endpoint: config.Endpoint{
				URL:      "/api/update",
				Type:     http.MethodPut,
				Response: map[string]any{"id": "uuid", "updated": true},
				Status:   200,
			},
			method:         http.MethodPut,
			expectedStatus: 200,
			description:    "Should handle PUT method",
		},
		{
			name: "PATCH_method",
			endpoint: config.Endpoint{
				URL:      "/api/patch",
				Type:     http.MethodPatch,
				Response: map[string]any{"id": "uuid", "patched": true},
				Status:   200,
			},
			method:         http.MethodPatch,
			expectedStatus: 200,
			description:    "Should handle PATCH method",
		},
		{
			name: "custom_status_code",
			endpoint: config.Endpoint{
				URL:      "/api/custom",
				Response: map[string]any{"message": "word"},
				Status:   418,
			},
			method:         http.MethodGet,
			expectedStatus: 418,
			description:    "Should return custom status code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{tt.endpoint},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(tt.method, tt.endpoint.URL, nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("%s: Expected status %d, got %d", tt.description, tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_PayloadValidationEdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       config.Endpoint
		payload        string
		expectedStatus int
		description    string
	}{
		{
			name: "nil_schema_field",
			endpoint: config.Endpoint{
				URL:      "/api/optional",
				Type:     http.MethodPost,
				Payload:  map[string]any{"required": "string", "optional": nil},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `{"required":"value"}`,
			expectedStatus: 200,
			description:    "Should allow missing optional fields (nil in schema)",
		},
		{
			name: "primitive_schema_validation",
			endpoint: config.Endpoint{
				URL:      "/api/primitive",
				Type:     http.MethodPost,
				Payload:  "string",
				Response: map[string]any{"ok": "word"},
			},
			payload:        `"any_value"`,
			expectedStatus: 200,
			description:    "Should accept any primitive when schema is primitive",
		},
		{
			name: "empty_array_schema",
			endpoint: config.Endpoint{
				URL:      "/api/empty_array",
				Type:     http.MethodPost,
				Payload:  []any{},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `[]`,
			expectedStatus: 200,
			description:    "Should accept empty array when schema is empty array",
		},
		{
			name: "empty_body_with_payload_schema",
			endpoint: config.Endpoint{
				URL:      "/api/empty_body",
				Type:     http.MethodPost,
				Payload:  map[string]any{"field": "string"},
				Response: map[string]any{"ok": "word"},
			},
			payload:        "",
			expectedStatus: 400,
			description:    "Should reject empty body when payload schema is provided",
		},
		{
			name: "DELETE_with_payload",
			endpoint: config.Endpoint{
				URL:      "/api/delete",
				Type:     http.MethodDelete,
				Payload:  map[string]any{"id": "string"},
				Response: map[string]any{"deleted": true},
			},
			payload:        `{"id":"123"}`,
			expectedStatus: 200,
			description:    "Should handle DELETE method with payload validation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{tt.endpoint},
			}
			handler := MakeHandler(cfg)

			var bodyReader io.Reader
			if tt.payload != "" {
				bodyReader = strings.NewReader(tt.payload)
			}
			req := httptest.NewRequest(tt.endpoint.Type, tt.endpoint.URL, bodyReader)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("%s: Expected status %d, got %d", tt.description, tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_ValidatePayloadStructureEdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       config.Endpoint
		payload        string
		expectedStatus int
		description    string
	}{
		{
			name: "nested_array_validation",
			endpoint: config.Endpoint{
				URL:      "/api/nested_array",
				Type:     http.MethodPost,
				Payload:  []any{[]any{map[string]any{"id": "string"}}},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `[[{"id":"value1"}],[{"id":"value2"}]]`,
			expectedStatus: 200,
			description:    "Should handle nested array validation",
		},
		{
			name: "nested_array_invalid_item",
			endpoint: config.Endpoint{
				URL:      "/api/nested_array_invalid",
				Type:     http.MethodPost,
				Payload:  []any{[]any{map[string]any{"id": "string"}}},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `[[{"id":"value1"}],[{"invalid":"value2"}]]`,
			expectedStatus: 400,
			description:    "Should reject invalid nested array items",
		},
		{
			name: "complex_nested_structure",
			endpoint: config.Endpoint{
				URL:  "/api/complex_nested",
				Type: http.MethodPost,
				Payload: map[string]any{
					"user": map[string]any{
						"profile": map[string]any{
							"settings": []any{map[string]any{"key": "string", "value": "string"}},
						},
					},
				},
				Response: map[string]any{"ok": "word"},
			},
			payload: `{
				"user": {
					"profile": {
						"settings": [
							{"key": "theme", "value": "dark"},
							{"key": "language", "value": "en"}
						]
					}
				}
			}`,
			expectedStatus: 200,
			description:    "Should handle complex nested structures",
		},
		{
			name: "primitive_schema_leaf_node",
			endpoint: config.Endpoint{
				URL:      "/api/primitive_leaf",
				Type:     http.MethodPost,
				Payload:  42,
				Response: map[string]any{"ok": "word"},
			},
			payload:        `"any_string_value"`,
			expectedStatus: 200,
			description:    "Should accept any primitive when schema is primitive (leaf node)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{tt.endpoint},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(tt.endpoint.Type, tt.endpoint.URL, strings.NewReader(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("%s: Expected status %d, got %d", tt.description, tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_ValidatePayloadStructureComplete(t *testing.T) {
	tests := []struct {
		name           string
		endpoint       config.Endpoint
		payload        string
		expectedStatus int
		description    string
	}{
		{
			name: "map_schema_with_non_map_body",
			endpoint: config.Endpoint{
				URL:      "/api/map_schema_non_map_body",
				Type:     http.MethodPost,
				Payload:  map[string]any{"field": "string"},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `"not_a_map"`,
			expectedStatus: 400,
			description:    "Should reject non-map body when schema expects map",
		},
		{
			name: "array_schema_with_non_array_body",
			endpoint: config.Endpoint{
				URL:      "/api/array_schema_non_array_body",
				Type:     http.MethodPost,
				Payload:  []any{map[string]any{"id": "string"}},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `{"id":"value"}`,
			expectedStatus: 400,
			description:    "Should reject non-array body when schema expects array",
		},
		{
			name: "map_schema_missing_required_field",
			endpoint: config.Endpoint{
				URL:      "/api/map_missing_field",
				Type:     http.MethodPost,
				Payload:  map[string]any{"required": "string", "optional": "string"},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `{"optional":"value"}`,
			expectedStatus: 400,
			description:    "Should reject when required field is missing",
		},
		{
			name: "nested_validation_failure",
			endpoint: config.Endpoint{
				URL:  "/api/nested_validation_failure",
				Type: http.MethodPost,
				Payload: map[string]any{
					"user": map[string]any{
						"profile": map[string]any{
							"name": "string",
						},
					},
				},
				Response: map[string]any{"ok": "word"},
			},
			payload: `{
				"user": {
					"profile": {
						"invalid_field": "value"
					}
				}
			}`,
			expectedStatus: 400,
			description:    "Should reject when nested validation fails",
		},
		{
			name: "array_item_validation_failure",
			endpoint: config.Endpoint{
				URL:      "/api/array_item_failure",
				Type:     http.MethodPost,
				Payload:  []any{map[string]any{"id": "string", "name": "string"}},
				Response: map[string]any{"ok": "word"},
			},
			payload:        `[{"id":"123","name":"John"},{"id":"456"}]`,
			expectedStatus: 400,
			description:    "Should reject when array item validation fails",
		},
		{
			name: "primitive_schema_accepts_any",
			endpoint: config.Endpoint{
				URL:      "/api/primitive_accepts_any",
				Type:     http.MethodPost,
				Payload:  "string",
				Response: map[string]any{"ok": "word"},
			},
			payload:        `123`,
			expectedStatus: 200,
			description:    "Should accept any primitive when schema is primitive",
		},
		{
			name: "primitive_schema_accepts_object",
			endpoint: config.Endpoint{
				URL:      "/api/primitive_accepts_object",
				Type:     http.MethodPost,
				Payload:  "string",
				Response: map[string]any{"ok": "word"},
			},
			payload:        `{"any": "object"}`,
			expectedStatus: 200,
			description:    "Should accept object when schema is primitive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{tt.endpoint},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(tt.endpoint.Type, tt.endpoint.URL, strings.NewReader(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("%s: Expected status %d, got %d", tt.description, tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_IndividualCachePerEndpoint(t *testing.T) {
	cacheSize1 := 2
	cacheSize2 := 3
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL: "/api/cached1",
				Response: map[string]any{
					"id":   "uuid",
					"data": "word",
				},
				Cache: &cacheSize1,
			},
			{
				URL: "/api/cached2",
				Response: map[string]any{
					"id":    "uuid",
					"value": "word",
				},
				Cache: &cacheSize2,
			},
			{
				URL: "/api/no-cache",
				Response: map[string]any{
					"id":     "uuid",
					"random": "word",
				},
				// Cache is nil - no caching
			},
		},
	}
	handler := MakeHandler(cfg)

	// Test first endpoint with cache
	req1 := httptest.NewRequest(http.MethodGet, "/api/cached1", nil)
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req1)

	if w1.Body.String() != w2.Body.String() {
		t.Error("Expected cached responses to be identical for /api/cached1")
	}

	// Test second endpoint with different cache
	req2 := httptest.NewRequest(http.MethodGet, "/api/cached2", nil)
	w3 := httptest.NewRecorder()
	handler.ServeHTTP(w3, req2)

	w4 := httptest.NewRecorder()
	handler.ServeHTTP(w4, req2)

	if w3.Body.String() != w4.Body.String() {
		t.Error("Expected cached responses to be identical for /api/cached2")
	}

	// Test endpoint without cache
	req3 := httptest.NewRequest(http.MethodGet, "/api/no-cache", nil)
	w5 := httptest.NewRecorder()
	handler.ServeHTTP(w5, req3)

	w6 := httptest.NewRecorder()
	handler.ServeHTTP(w6, req3)

	// Responses should be different (no caching)
	if w5.Body.String() == w6.Body.String() {
		t.Error("Expected different responses for /api/no-cache (no caching)")
	}
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}
