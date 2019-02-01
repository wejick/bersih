package service

import (
	"context"

	"github.com/wejick/bersih/pkg/autocomplete/model"
)

// ProfileDataList list of profile
type ProfileDataList struct {
	TotalData int64
	Profile   []ProfileData
}

// ProfileData data block of profile
type ProfileData struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// GetProfile returns list of profile
func (S *Service) GetProfile(ctx context.Context) (profileDataList ProfileDataList, err error) {
	profileList, err := S.repo.GetProfile(ctx)
	if err != nil {
		return
	}

	for _, profile := range profileList.Data {
		profileDataList.Profile = append(profileDataList.Profile, profileToProfileData(profile))
	}
	profileDataList.TotalData = profileList.TotalData

	return
}

func profileDataToModel(profileData ProfileData) (profile model.Profile) {
	profile = model.Profile{
		Avatar: profileData.Avatar,
		Name:   profileData.Name,
	}

	return
}

func profileToProfileData(profile model.Profile) (profileData ProfileData) {
	profileData = ProfileData{
		Avatar: profile.Avatar,
		Name:   profile.Name,
	}
	return
}
