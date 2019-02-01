package http

import (
	"reflect"
	"testing"

	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
	autocompleteService "github.com/wejick/bersih/pkg/autocomplete/service"
)

func TestNew(t *testing.T) {
	type args struct {
		service *autocompleteService.Service
	}
	tests := []struct {
		name        string
		args        args
		wantHandler *AutocompleteHTTPHandler
	}{
		{
			name:        "empty",
			wantHandler: &AutocompleteHTTPHandler{},
		},
		{
			name: "empty",
			args: args{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{}),
			},
			wantHandler: &AutocompleteHTTPHandler{
				service: autocompleteService.New(&autocompleteRepo.RepoMock{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHandler := New(tt.args.service); !reflect.DeepEqual(gotHandler, tt.wantHandler) {
				t.Errorf("New() = %v, want %v", gotHandler, tt.wantHandler)
			}
		})
	}
}
