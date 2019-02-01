package http

import (
	"net/http"

	"github.com/wejick/bersih/pkg/autocomplete/service"
	"github.com/wejick/bersih/pkg/httputil"

	"github.com/julienschmidt/httprouter"
)

// ProfileResponse is http response for profile
type ProfileResponse struct {
	Header Header                `json:"header"`
	Data   []service.ProfileData `json:"data"`
}

// GetProfile returns profile list
func (H *AutocompleteHTTPHandler) GetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := H.service.GetProfile(r.Context())
	if err != nil {
		httputil.ResponseError("couldn't get profile", http.StatusInternalServerError, w)
		return
	}

	response := ProfileResponse{
		Header: Header{
			TotalData: data.TotalData,
		},
		Data: data.Profile,
	}

	httputil.ResponseJSON(response, http.StatusOK, w)
}
