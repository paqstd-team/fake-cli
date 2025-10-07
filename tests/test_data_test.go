package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
	"github.com/paqstd-team/fake-cli/handler"
)

func TestData_GenerateAllFieldsSingle(t *testing.T) {
	fields := map[string]any{
		// id
		"uuid": "uuid",
		// geography
		"city": "city", "state": "state", "country": "country", "latitude": "latitude", "longitude": "longitude",
		// person
		"name": "name", "name_prefix": "name_prefix", "name_suffix": "name_suffix", "first_name": "first_name", "last_name": "last_name",
		"gender": "gender", "ssn": "ssn", "hobby": "hobby", "email": "email", "phone": "phone", "username": "username", "password": "password",
		// text
		"paragraph": "paragraph", "sentence": "sentence", "phrase": "phrase", "quote": "quote", "word": "word",
		// date/time
		"date": "date", "second": "second", "minute": "minute", "hour": "hour", "month": "month", "day": "day", "year": "year",
		// internet
		"url": "url", "domain": "domain", "ip": "ip",
		// numbers
		"int": "int", "float": "float",
		// default passthrough
		"unknown": "unknown",
		// nested map and array to exercise recursion
		"nested": map[string]any{"inner": "word"},
		"arr":    []any{"sentence", map[string]any{"deep": "city"}},
	}
	cfg := config.Config{Endpoints: []config.Endpoint{{URL: "/data", Fields: fields, Response: "single"}}}
	h := handler.MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/data", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}

	var obj map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &obj); err != nil {
		t.Fatalf("json: %v", err)
	}
	keys := []string{"uuid", "city", "state", "country", "latitude", "longitude", "name", "name_prefix", "name_suffix", "first_name", "last_name", "gender", "ssn", "hobby", "email", "phone", "username", "password", "paragraph", "sentence", "phrase", "quote", "word", "date", "second", "minute", "hour", "month", "day", "year", "url", "domain", "ip", "int", "float", "unknown", "nested", "arr"}
	for _, k := range keys {
		if _, ok := obj[k]; !ok {
			t.Fatalf("missing key %s", k)
		}
	}
}

func TestData_MapStringStringPath(t *testing.T) {
	// Exercise generateData(map[string]string) path
	fields := map[string]string{"id": "uuid", "city": "city"}
	cfg := config.Config{Endpoints: []config.Endpoint{{URL: "/mss", Fields: fields, Response: "single"}}}
	h := handler.MakeHandler(cfg)
	req := httptest.NewRequest(http.MethodGet, "/mss", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: %d", w.Code)
	}
	var obj map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &obj); err != nil {
		t.Fatalf("json: %v", err)
	}
	if _, ok := obj["id"]; !ok {
		t.Fatalf("missing id")
	}
}
