package service

import (
	"reflect"
	"testing"

	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
)

func TestNew(t *testing.T) {
	repoMock := &autocompleteRepo.RepoMock{}
	type args struct {
		repo autocompleteRepo.Repo
	}
	tests := []struct {
		name        string
		args        args
		wantService *Service
	}{
		{
			name:        "empty",
			wantService: &Service{},
		},
		{
			name: "with mock",
			args: args{
				repo: repoMock,
			},
			wantService: &Service{
				repo: repoMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotService := New(tt.args.repo); !reflect.DeepEqual(gotService, tt.wantService) {
				t.Errorf("New() = %v, want %v", gotService, tt.wantService)
			}
		})
	}
}
