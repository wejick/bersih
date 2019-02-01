package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/wejick/bersih/pkg/autocomplete/model"
	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
)

func Test_textToTextData(t *testing.T) {
	type args struct {
		text model.Text
	}
	tests := []struct {
		name         string
		args         args
		wantTextData TextData
	}{
		{
			name: "empty",
		},
		{
			name: "full",
			args: args{
				text: model.Text{
					Data: "data",
					URL:  "http://example.com",
				},
			},
			wantTextData: TextData{
				Data: "data",
				URL:  "http://example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTextData := textToTextData(tt.args.text); !reflect.DeepEqual(gotTextData, tt.wantTextData) {
				t.Errorf("textToTextData() = %v, want %v", gotTextData, tt.wantTextData)
			}
		})
	}
}

func Test_textDataToModel(t *testing.T) {
	type args struct {
		textData TextData
	}
	tests := []struct {
		name     string
		args     args
		wantText model.Text
	}{
		{
			name: "empty",
		},
		{
			name: "full",
			args: args{
				textData: TextData{
					Data: "data",
					URL:  "http://example.com"},
			},
			wantText: model.Text{
				Data: "data",
				URL:  "http://example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotText := textDataToModel(tt.args.textData); !reflect.DeepEqual(gotText, tt.wantText) {
				t.Errorf("textDataToModel() = %v, want %v", gotText, tt.wantText)
			}
		})
	}
}

func TestService_GetText(t *testing.T) {
	type fields struct {
		repo autocompleteRepo.Repo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantTextDataList TextDataList
		wantErr          bool
	}{
		{
			name: "empty",
			fields: fields{
				repo: &autocompleteRepo.RepoMock{
					GetTextFunc: func(ctx context.Context) (text autocompleteRepo.TextList, err error) {
						return
					},
				},
			},
		},
		{
			name: "with data",
			fields: fields{
				repo: &autocompleteRepo.RepoMock{
					GetTextFunc: func(ctx context.Context) (text autocompleteRepo.TextList, err error) {
						text = autocompleteRepo.TextList{
							TotalData: 2,
							Data: []model.Text{model.Text{
								Data: "data1",
								URL:  "example.com",
							},
								model.Text{
									Data: "data2",
									URL:  "example.com",
								},
							},
						}
						return
					},
				},
			},
			wantTextDataList: TextDataList{
				TotalData: 2,
				Text:      []TextData{TextData{Data: "data1", URL: "example.com"}, TextData{Data: "data2", URL: "example.com"}},
			},
		},
		{
			name: "with error",
			fields: fields{
				repo: &autocompleteRepo.RepoMock{
					GetTextFunc: func(ctx context.Context) (text autocompleteRepo.TextList, err error) {
						err = fmt.Errorf("error")
						return
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			S := &Service{
				repo: tt.fields.repo,
			}
			gotTextDataList, err := S.GetText(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTextDataList, tt.wantTextDataList) {
				t.Errorf("Service.GetText() = %v, want %v", gotTextDataList, tt.wantTextDataList)
			}
		})
	}
}
