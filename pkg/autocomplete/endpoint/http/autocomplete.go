package http

import (
	autocompleteService "github.com/wejick/bersih/pkg/autocomplete/service"
)

// Header http header
type Header struct {
	TotalData int64 `json:"total_data"`
}

// AutocompleteHTTPHandler wrap the handler
type AutocompleteHTTPHandler struct {
	service *autocompleteService.Service
}

// New create handler instance
func New(service *autocompleteService.Service) (handler *AutocompleteHTTPHandler) {
	handler = &AutocompleteHTTPHandler{
		service: service,
	}

	return
}
