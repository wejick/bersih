package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/wejick/bersih/pkg/autocomplete/model"
	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
	autocompleteService "github.com/wejick/bersih/pkg/autocomplete/service"
)

func TestAutocompleteHTTPHandler_GetProfile(t *testing.T) {
	type fields struct {
		service *autocompleteService.Service
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		wantCode            int
		wantProfileResponse ProfileResponse
	}{
		{
			name: "empty",
			fields: fields{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{
					GetProfileFunc: func(ctx context.Context) (profile autocompleteRepo.ProfileList, err error) {
						return
					},
				}),
			},
			wantCode: http.StatusOK,
		},
		{
			name: "with data",
			fields: fields{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{
					GetProfileFunc: func(ctx context.Context) (profile autocompleteRepo.ProfileList, err error) {
						profile = autocompleteRepo.ProfileList{
							TotalData: 2,
							Data:      []model.Profile{model.Profile{}, model.Profile{}},
						}
						return
					},
				}),
			},
			wantCode: http.StatusOK,
			wantProfileResponse: ProfileResponse{
				Header: Header{TotalData: 2},
				Data:   []autocompleteService.ProfileData{autocompleteService.ProfileData{}, autocompleteService.ProfileData{}}},
		},
		{
			name: "with error",
			fields: fields{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{
					GetProfileFunc: func(ctx context.Context) (profile autocompleteRepo.ProfileList, err error) {
						err = fmt.Errorf("error")
						return
					},
				}),
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			H := &AutocompleteHTTPHandler{
				service: tt.fields.service,
			}
			router := httprouter.New()
			router.GET("/", H.GetProfile)
			recorder := httptest.NewRecorder()
			emptyGetRequest, _ := http.NewRequest("GET", "/", nil)
			router.ServeHTTP(recorder, emptyGetRequest)
			assert.Equal(t, tt.wantCode, recorder.Code, "error code")

			var response ProfileResponse
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err, "unmarshall no error")
			assert.Equal(t, tt.wantProfileResponse, response, "handler response")
		})
	}
}
