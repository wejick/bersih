package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wejick/bersih/pkg/autocomplete/model"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
	autocompleteService "github.com/wejick/bersih/pkg/autocomplete/service"
)

func TestAutocompleteHTTPHandler_GetText(t *testing.T) {
	type fields struct {
		service *autocompleteService.Service
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantCode         int
		wantTextResponse TextResponse
	}{
		{
			name: "empty ",
			fields: fields{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{
					GetTextFunc: func(ctx context.Context) (text autocompleteRepo.TextList, err error) {
						return
					},
				}),
			},
			wantCode: http.StatusOK,
		},
		{
			name: "with data ",
			fields: fields{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{
					GetTextFunc: func(ctx context.Context) (text autocompleteRepo.TextList, err error) {
						text = autocompleteRepo.TextList{
							TotalData: 2,
							Data:      []model.Text{model.Text{}, model.Text{}},
						}
						return
					},
				}),
			},
			wantCode: http.StatusOK,
			wantTextResponse: TextResponse{
				Header: Header{TotalData: 2},
				Data:   []autocompleteService.TextData{autocompleteService.TextData{}, autocompleteService.TextData{}},
			},
		},
		{
			name: "with error",
			fields: fields{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{
					GetTextFunc: func(ctx context.Context) (text autocompleteRepo.TextList, err error) {
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
			router.GET("/", H.GetText)
			recorder := httptest.NewRecorder()
			emptyGetRequest, _ := http.NewRequest("GET", "/", nil)
			router.ServeHTTP(recorder, emptyGetRequest)
			assert.Equal(t, tt.wantCode, recorder.Code)

			var response TextResponse
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err, "unmarshall no error")
			assert.Equal(t, tt.wantTextResponse, response, "handler response")
		})
	}
}
