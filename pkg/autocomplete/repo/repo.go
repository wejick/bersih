package repo

import (
	"context"

	"github.com/wejick/bersih/pkg/autocomplete/model"
)

//go:generate moq -out repo_moq.go . Repo

// TextList contains texts data
type TextList struct {
	TotalData int64
	Data      []model.Text
}

// ProfileList contains profiles data
type ProfileList struct {
	TotalData int64
	Data      []model.Profile
}

// Repo repository for autocomplete
type Repo interface {
	CreateText(context.Context, model.Text) error
	GetText(context.Context) (TextList, error)
	UpdateText(context.Context, model.Text) error
	DeleteText(context.Context, model.Text) error

	CreateProfile(context.Context, model.Profile) error
	GetProfile(context.Context) (ProfileList, error)
	UpdateProfile(context.Context, model.Profile) error
	DeleteProfile(context.Context, model.Profile) error

	Initialize() error
}
