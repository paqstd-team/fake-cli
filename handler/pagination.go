package handler

import (
	"net/http"
	"strconv"
)

func getPaginationParams(r *http.Request) (page int, perPage int) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		page = 1 // Set default page to 1
	} else {
		page, _ = strconv.Atoi(pageStr)
	}

	perPageStr := r.URL.Query().Get("per_page")
	if perPageStr == "" {
		perPage = 10 // Set default per_page to 10
	} else {
		perPage, _ = strconv.Atoi(perPageStr)
	}

	return page, perPage
}
