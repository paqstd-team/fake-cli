package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
)

func TestPagination_Parameters(t *testing.T) {
	tests := []struct {
		name        string
		page        string
		perPage     string
		expectedLen int
		description string
	}{
		{
			name:        "default_pagination",
			page:        "",
			perPage:     "",
			expectedLen: 10,
			description: "Should use default values when no parameters provided",
		},
		{
			name:        "custom_page_size",
			page:        "1",
			perPage:     "5",
			expectedLen: 5,
			description: "Should respect custom per_page parameter",
		},
		{
			name:        "large_page_size",
			page:        "1",
			perPage:     "50",
			expectedLen: 50,
			description: "Should handle large page sizes",
		},
		{
			name:        "page_zero",
			page:        "0",
			perPage:     "10",
			expectedLen: 0,
			description: "Should return empty list when page is 0",
		},
		{
			name:        "negative_page",
			page:        "-1",
			perPage:     "10",
			expectedLen: 0,
			description: "Should return empty list when page is negative",
		},
		{
			name:        "zero_per_page",
			page:        "1",
			perPage:     "0",
			expectedLen: 0,
			description: "Should return empty list when per_page is 0",
		},
		{
			name:        "negative_per_page",
			page:        "1",
			perPage:     "-5",
			expectedLen: 0,
			description: "Should return empty list when per_page is negative",
		},
		{
			name:        "invalid_page_string",
			page:        "invalid",
			perPage:     "10",
			expectedLen: 0,
			description: "Should handle invalid page parameter gracefully",
		},
		{
			name:        "invalid_per_page_string",
			page:        "1",
			perPage:     "invalid",
			expectedLen: 0,
			description: "Should handle invalid per_page parameter gracefully",
		},
		{
			name:        "very_large_page",
			page:        "999999",
			perPage:     "10",
			expectedLen: 10,
			description: "Should return per_page items for any valid page number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{
					{
						URL: "/api/products",
						Response: []any{
							map[string]any{"id": "uuid", "name": "word", "price": "float"},
						},
					},
				},
			}
			handler := MakeHandler(cfg)

			u := &url.URL{Path: "/api/products"}
			q := u.Query()
			if tt.page != "" {
				q.Set("page", tt.page)
			}
			if tt.perPage != "" {
				q.Set("per_page", tt.perPage)
			}
			u.RawQuery = q.Encode()

			req := httptest.NewRequest(http.MethodGet, u.String(), nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var items []any
			if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if len(items) != tt.expectedLen {
				t.Errorf("%s: Expected %d items, got %d", tt.description, tt.expectedLen, len(items))
			}
		})
	}
}

func TestPagination_EmptyTemplate(t *testing.T) {
	tests := []struct {
		name        string
		template    []any
		page        string
		perPage     string
		expectedLen int
	}{
		{
			name:        "empty_template_default",
			template:    []any{},
			page:        "",
			perPage:     "",
			expectedLen: 10,
		},
		{
			name:        "empty_template_custom_size",
			template:    []any{},
			page:        "1",
			perPage:     "7",
			expectedLen: 7,
		},
		{
			name:        "empty_template_page_zero",
			template:    []any{},
			page:        "0",
			perPage:     "10",
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{
					{
						URL:      "/api/empty",
						Response: tt.template,
					},
				},
			}
			handler := MakeHandler(cfg)

			u := &url.URL{Path: "/api/empty"}
			q := u.Query()
			if tt.page != "" {
				q.Set("page", tt.page)
			}
			if tt.perPage != "" {
				q.Set("per_page", tt.perPage)
			}
			u.RawQuery = q.Encode()

			req := httptest.NewRequest(http.MethodGet, u.String(), nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var items []any
			if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if len(items) != tt.expectedLen {
				t.Errorf("Expected %d items, got %d", tt.expectedLen, len(items))
			}
		})
	}
}

func TestPagination_ComplexData(t *testing.T) {
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL: "/api/orders",
				Response: []any{
					map[string]any{
						"order_id": "uuid",
						"customer": map[string]any{
							"id":    "uuid",
							"name":  "name",
							"email": "email",
						},
						"items": []any{
							map[string]any{
								"product_id": "uuid",
								"quantity":   "int",
								"price":      "float",
							},
						},
						"total":      "float",
						"status":     "word",
						"created_at": "date",
					},
				},
			},
		},
	}
	handler := MakeHandler(cfg)

	tests := []struct {
		name        string
		page        string
		perPage     string
		expectedLen int
	}{
		{"page_1_size_3", "1", "3", 3},
		{"page_2_size_3", "2", "3", 3},
		{"page_3_size_3", "3", "3", 3},
		{"page_4_size_3", "4", "3", 3},
		{"page_5_size_3", "5", "3", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &url.URL{Path: "/api/orders"}
			q := u.Query()
			q.Set("page", tt.page)
			q.Set("per_page", tt.perPage)
			u.RawQuery = q.Encode()

			req := httptest.NewRequest(http.MethodGet, u.String(), nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var items []any
			if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if len(items) != tt.expectedLen {
				t.Errorf("Expected %d items, got %d", tt.expectedLen, len(items))
			}

			for i, item := range items {
				order, ok := item.(map[string]any)
				if !ok {
					t.Errorf("Item %d is not a map", i)
					continue
				}

				if _, exists := order["order_id"]; !exists {
					t.Errorf("Item %d missing order_id", i)
				}
				if _, exists := order["customer"]; !exists {
					t.Errorf("Item %d missing customer", i)
				}
				if _, exists := order["items"]; !exists {
					t.Errorf("Item %d missing items", i)
				}
			}
		})
	}
}

func TestPagination_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		page        string
		perPage     string
		expectedLen int
		description string
	}{
		{
			name:        "float_page",
			page:        "1.5",
			perPage:     "10",
			expectedLen: 0,
			description: "Should handle float page numbers",
		},
		{
			name:        "float_per_page",
			page:        "1",
			perPage:     "5.7",
			expectedLen: 0,
			description: "Should handle float per_page numbers",
		},
		{
			name:        "very_large_per_page",
			page:        "1",
			perPage:     "999999",
			expectedLen: 999999,
			description: "Should handle very large per_page values",
		},
		{
			name:        "zero_page_with_per_page",
			page:        "0",
			perPage:     "5",
			expectedLen: 0,
			description: "Should return empty for page 0 regardless of per_page",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{
					{
						URL: "/api/test",
						Response: []any{
							map[string]any{"id": "uuid", "value": "word"},
						},
					},
				},
			}
			handler := MakeHandler(cfg)

			u := &url.URL{Path: "/api/test"}
			q := u.Query()
			q.Set("page", tt.page)
			q.Set("per_page", tt.perPage)
			u.RawQuery = q.Encode()

			req := httptest.NewRequest(http.MethodGet, u.String(), nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var items []any
			if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if len(items) != tt.expectedLen {
				t.Errorf("%s: Expected %d items, got %d", tt.description, tt.expectedLen, len(items))
			}
		})
	}
}
