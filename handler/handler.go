package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeHandler(config Config) http.Handler {
	mux := http.NewServeMux()

	for _, endpoint := range config.Endpoints {
		handlerFunc := makeHandlerFunc(endpoint.Fields, endpoint.Response)
		mux.HandleFunc(endpoint.URL, handlerFunc)
	}

	return mux
}

func makeHandlerFunc(fields map[string]string, responseType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
