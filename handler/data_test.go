package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
)

func TestData_GenerateFields(t *testing.T) {
	tests := []struct {
		name         string
		fields       map[string]any
		expectedKeys []string
	}{
		{
			name: "user_profile_fields",
			fields: map[string]any{
				"id":         "uuid",
				"username":   "username",
				"email":      "email",
				"first_name": "first_name",
				"last_name":  "last_name",
				"phone":      "phone",
				"created_at": "date",
			},
			expectedKeys: []string{"id", "username", "email", "first_name", "last_name", "phone", "created_at"},
		},
		{
			name: "product_catalog_fields",
			fields: map[string]any{
				"product_id":  "uuid",
				"title":       "word",
				"description": "sentence",
				"price":       "float",
				"category":    "word",
				"in_stock":    "bool",
				"tags":        []any{"word", "word", "word"},
			},
			expectedKeys: []string{"product_id", "title", "description", "price", "category", "in_stock", "tags"},
		},
		{
			name: "geographic_data_fields",
			fields: map[string]any{
				"location_id": "uuid",
				"city":        "city",
				"state":       "state",
				"country":     "country",
				"latitude":    "latitude",
				"longitude":   "longitude",
				"postal_code": "zip",
			},
			expectedKeys: []string{"location_id", "city", "state", "country", "latitude", "longitude", "postal_code"},
		},
		{
			name: "nested_structure_fields",
			fields: map[string]any{
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
				"total":  "float",
				"status": "word",
			},
			expectedKeys: []string{"order_id", "customer", "items", "total", "status"},
		},
		{
			name: "mixed_data_types",
			fields: map[string]any{
				"id":            "uuid",
				"text_content":  "paragraph",
				"short_text":    "word",
				"numeric_value": "int",
				"decimal_value": "float",
				"boolean_flag":  "bool",
				"url_link":      "url",
				"ip_address":    "ip",
				"domain_name":   "domain",
			},
			expectedKeys: []string{"id", "text_content", "short_text", "numeric_value", "decimal_value", "boolean_flag", "url_link", "ip_address", "domain_name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{
					{
						URL:      "/api/data",
						Response: tt.fields,
					},
				},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(http.MethodGet, "/api/data", nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var response map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for _, key := range tt.expectedKeys {
				if _, exists := response[key]; !exists {
					t.Errorf("Expected key '%s' not found in response", key)
				}
			}

			if len(response) != len(tt.expectedKeys) {
				t.Errorf("Expected %d keys, got %d", len(tt.expectedKeys), len(response))
			}
		})
	}
}

func TestData_MapStringStringPath(t *testing.T) {
	tests := []struct {
		name         string
		fields       map[string]string
		expectedKeys []string
	}{
		{
			name: "simple_string_map",
			fields: map[string]string{
				"id":    "uuid",
				"name":  "name",
				"email": "email",
			},
			expectedKeys: []string{"id", "name", "email"},
		},
		{
			name: "location_string_map",
			fields: map[string]string{
				"city":    "city",
				"state":   "state",
				"country": "country",
			},
			expectedKeys: []string{"city", "state", "country"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{
					{
						URL:      "/api/simple",
						Response: tt.fields,
					},
				},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(http.MethodGet, "/api/simple", nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var response map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for _, key := range tt.expectedKeys {
				if _, exists := response[key]; !exists {
					t.Errorf("Expected key '%s' not found in response", key)
				}
			}
		})
	}
}

func TestData_AllFieldTypes(t *testing.T) {
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL: "/api/all-fields",
				Response: map[string]any{
					"uuid":          "uuid",
					"city":          "city",
					"state":         "state",
					"country":       "country",
					"latitude":      "latitude",
					"longitude":     "longitude",
					"name":          "name",
					"name_prefix":   "name_prefix",
					"name_suffix":   "name_suffix",
					"first_name":    "first_name",
					"last_name":     "last_name",
					"gender":        "gender",
					"ssn":           "ssn",
					"hobby":         "hobby",
					"email":         "email",
					"phone":         "phone",
					"username":      "username",
					"password":      "password",
					"paragraph":     "paragraph",
					"sentence":      "sentence",
					"phrase":        "phrase",
					"quote":         "quote",
					"word":          "word",
					"date":          "date",
					"second":        "second",
					"minute":        "minute",
					"hour":          "hour",
					"month":         "month",
					"day":           "day",
					"year":          "year",
					"url":           "url",
					"domain":        "domain",
					"ip":            "ip",
					"int":           "int",
					"float":         "float",
					"unknown_field": "unknown_value",
				},
			},
		},
	}
	handler := MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/all-fields", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedFields := []string{
		"uuid", "city", "state", "country", "latitude", "longitude",
		"name", "name_prefix", "name_suffix", "first_name", "last_name",
		"gender", "ssn", "hobby", "email", "phone", "username", "password",
		"paragraph", "sentence", "phrase", "quote", "word",
		"date", "second", "minute", "hour", "month", "day", "year",
		"url", "domain", "ip", "int", "float", "unknown_field",
	}

	for _, field := range expectedFields {
		if _, exists := response[field]; !exists {
			t.Errorf("Expected field '%s' not found in response", field)
		}
	}

	if response["unknown_field"] != "unknown_value" {
		t.Errorf("Expected unknown_field to be 'unknown_value', got %v", response["unknown_field"])
	}
}

func TestData_ComplexNestedStructures(t *testing.T) {
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL: "/api/complex",
				Response: map[string]any{
					"company": map[string]any{
						"id":   "uuid",
						"name": "company",
						"address": map[string]any{
							"street": "street",
							"city":   "city",
							"zip":    "zip",
						},
						"employees": []any{
							map[string]any{
								"id":       "uuid",
								"name":     "name",
								"position": "job_title",
							},
						},
					},
					"metadata": map[string]any{
						"created_at": "date",
						"version":    "int",
						"active":     "bool",
					},
				},
			},
		},
	}
	handler := MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/complex", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if _, exists := response["company"]; !exists {
		t.Error("Expected 'company' key not found")
	}

	if _, exists := response["metadata"]; !exists {
		t.Error("Expected 'metadata' key not found")
	}

	company, ok := response["company"].(map[string]any)
	if !ok {
		t.Fatal("Expected 'company' to be a map")
	}

	if _, exists := company["address"]; !exists {
		t.Error("Expected 'address' key not found in company")
	}

	if _, exists := company["employees"]; !exists {
		t.Error("Expected 'employees' key not found in company")
	}
}
