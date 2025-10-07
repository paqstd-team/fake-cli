package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paqstd-team/fake-cli/v2/cache"
	"github.com/paqstd-team/fake-cli/v2/config"
)

// JSONMarshal is used to marshal response data. It is a variable to allow tests
// to inject failures and exercise error-handling branches.
var JSONMarshal = json.Marshal

func MakeHandler(config config.Config) http.Handler {
	mux := mux.NewRouter()

	for _, endpoint := range config.Endpoints {
		method := endpoint.Type
		if method == "" {
			method = http.MethodGet
		}

		// Create individual cache for this endpoint if cache is specified
		var endpointCache *cache.Cache
		if endpoint.Cache != nil {
			endpointCache = cache.NewCache(*endpoint.Cache)
		}

		handlerFunc := makeHandlerFunc(endpoint.Response, endpoint.Status, endpoint.Payload, endpointCache, method)
		mux.HandleFunc(endpoint.URL, handlerFunc).Methods(method)
	}

	return mux
}

func makeHandlerFunc(response any, status int, payload any, cache *cache.Cache, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if status == 0 {
			status = http.StatusOK
		}

		// Validate payload when schema is provided and method commonly carries a body
		if payload != nil && (method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch || method == http.MethodDelete) {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
				return
			}
			// Empty body is invalid when a payload schema is specified
			if len(bodyBytes) == 0 {
				http.Error(w, "Empty body", http.StatusBadRequest)
				return
			}
			var body any
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				http.Error(w, "Invalid JSON body", http.StatusBadRequest)
				return
			}
			if !validatePayloadStructure(payload, body) {
				http.Error(w, "Payload does not match schema", http.StatusBadRequest)
				return
			}
		}

		// Only use cache for GET requests and if cache is configured
		var cacheKey string
		if r.Method == http.MethodGet && cache != nil {
			cacheKey = r.Method + ":" + r.URL.Path + r.URL.RawQuery
			cacheValue, cacheHit := cache.Get(cacheKey)
			if cacheHit {
				w.WriteHeader(status)
				if status != http.StatusNoContent {
					w.Write([]byte(cacheValue.(string)))
				}
				return
			}
		}

		var data interface{}
		switch v := response.(type) {
		case []interface{}:
			// Treat top-level array as a list: use the first element as template
			var template any
			if len(v) > 0 {
				template = v[0]
			} else {
				template = map[string]any{}
			}
			page, perPage := getPaginationParams(r)
			data = generateDataList(template, page, perPage)
		default:
			// Treat maps and primitives as a single object response
			data = generateData(response)
		}

		jsonData, err := JSONMarshal(data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating JSON: %v", err), http.StatusInternalServerError)
			return
		}

		// Write status first
		w.WriteHeader(status)
		if status == http.StatusNoContent {
			return
		}
		// Cache only GET responses and if cache is configured
		if cacheKey != "" && cache != nil {
			cache.Set(cacheKey, string(jsonData))
		}
		w.Write(jsonData)
	}
}

// validatePayloadStructure ensures that the given body matches the shape of the schema.
// It validates structure only (objects vs arrays and required keys), not the primitive value types.
func validatePayloadStructure(schema any, body any) bool {
	switch s := schema.(type) {
	case map[string]any:
		bmap, ok := body.(map[string]any)
		if !ok {
			return false
		}
		for k, sub := range s {
			if sub == nil {
				// nil in schema means optional; skip
				continue
			}
			if _, exists := bmap[k]; !exists {
				return false
			}
			if !validatePayloadStructure(sub, bmap[k]) {
				return false
			}
		}
		return true
	case []any:
		// Treat first element as template if present
		var template any
		if len(s) > 0 {
			template = s[0]
		} else {
			template = map[string]any{}
		}
		barr, ok := body.([]any)
		if !ok {
			return false
		}
		for _, item := range barr {
			if !validatePayloadStructure(template, item) {
				return false
			}
		}
		return true
	default:
		// For leaf nodes (including strings in schema), accept any primitive
		return true
	}
}
