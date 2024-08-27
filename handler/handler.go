package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paqstd-team/fake-cli/cache"
	"github.com/paqstd-team/fake-cli/config"
)

func MakeHandler(config config.Config) http.Handler {
	mux := mux.NewRouter()
	cache := cache.NewCache(config.Cache)

	for _, endpoint := range config.Endpoints {
		handlerFunc := makeHandlerFunc(endpoint.Fields, endpoint.Response, cache)
		mux.HandleFunc(endpoint.URL, handlerFunc)
	}

	return mux
}

func makeHandlerFunc(fields any, responseType string, cache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cacheKey := r.URL.Path + r.URL.RawQuery
		cacheValue, cacheHit := cache.Get(cacheKey)
		if cacheHit {
			w.Write([]byte(cacheValue.(string)))
			return
		}

		var data interface{}
		if responseType == "list" {
			page, perPage := getPaginationParams(r)
			data = generateDataList(fields, page, perPage)
		} else {
			data = generateData(fields)
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating JSON: %v", err), http.StatusInternalServerError)
			return
		}

		cache.Set(cacheKey, string(jsonData))
		w.Write(jsonData)
	}
}
