package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/wejick/bersih/pkg/autocomplete/model"
	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
)

func Test_profileDataToModel(t *testing.T) {
	type args struct {
		profileData ProfileData
	}
	tests := []struct {
		name        string
		args        args
		wantProfile model.Profile
	}{
		{
			name: "empty",
		},
		{
			name: "full",
			args: args{
				profileData: ProfileData{
					Avatar: "url_to_avatar",
					Name:   "profilename",
				},
			},
			wantProfile: model.Profile{
				Avatar: "url_to_avatar",
				Name:   "profilename",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProfile := profileDataToModel(tt.args.profileData); !reflect.DeepEqual(gotProfile, tt.wantProfile) {
				t.Errorf("profileDataToModel() = %v, want %v", gotProfile, tt.wantProfile)
			}
		})
	}
}

func Test_profileToProfileData(t *testing.T) {
	type args struct {
		profile model.Profile
	}
	tests := []struct {
		name            string
		args            args
		wantProfileData ProfileData
	}{
		{
			name: "empty",
		},
		{
			name: "full",
			args: args{
				profile: model.Profile{
					Avatar: "url_to_avatar",
					Name:   "profilename",
				},
			},
			wantProfileData: ProfileData{
				Avatar: "url_to_avatar",
				Name:   "profilename",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProfileData := profileToProfileData(tt.args.profile); !reflect.DeepEqual(gotProfileData, tt.wantProfileData) {
				t.Errorf("profileToProfileData() = %v, want %v", gotProfileData, tt.wantProfileData)
			}
		})
	}
}

func TestService_GetProfile(t *testing.T) {
	type fields struct {
		repo autocompleteRepo.Repo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		wantProfileDataList ProfileDataList
		wantErr             bool
	}{
		{
			name: "empty",
			fields: fields{
				repo: &autocompleteRepo.RepoMock{
					GetProfileFunc: func(ctx context.Context) (repo autocompleteRepo.ProfileList, err error) {
						return
					},
				},
			},
		},
		{
			name: "with data",
			fields: fields{
				repo: &autocompleteRepo.RepoMock{
					GetProfileFunc: func(ctx context.Context) (repo autocompleteRepo.ProfileList, err error) {
						repo = autocompleteRepo.ProfileList{
							TotalData: 2,
							Data:      []model.Profile{model.Profile{Name: "name1"}, model.Profile{Name: "name2"}},
						}
						return
					},
				},
			},
			wantProfileDataList: ProfileDataList{
				TotalData: 2,
				Profile:   []ProfileData{ProfileData{Name: "name1"}, ProfileData{Name: "name2"}},
			},
		},
		{
			name: "with error",
			fields: fields{
				repo: &autocompleteRepo.RepoMock{
					GetProfileFunc: func(ctx context.Context) (repo autocompleteRepo.ProfileList, err error) {
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
			gotProfileDataList, err := S.GetProfile(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProfileDataList, tt.wantProfileDataList) {
				t.Errorf("Service.GetProfile() = %v, want %v", gotProfileDataList, tt.wantProfileDataList)
			}
		})
	}
}
