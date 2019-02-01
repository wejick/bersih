package http

import (
	"net/http"

	"github.com/wejick/bersih/pkg/autocomplete/service"
	"github.com/wejick/bersih/pkg/httputil"

	"github.com/julienschmidt/httprouter"
)

// TextResponse is http response for text
type TextResponse struct {
	Header Header             `json:"header"`
	Data   []service.TextData `json:"data"`
}

// GetText returns text list
func (H *AutocompleteHTTPHandler) GetText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := H.service.GetText(r.Context())
	if err != nil {
		httputil.ResponseError("couldn't get text", http.StatusInternalServerError, w)
		return
	}

	response := TextResponse{
		Header: Header{
			TotalData: data.TotalData,
		},
		Data: data.Text,
	}

	httputil.ResponseJSON(response, http.StatusOK, w)
}
