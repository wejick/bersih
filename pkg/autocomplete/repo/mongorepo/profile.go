package mongorepo

import (
	"context"

	"github.com/wejick/bersih/pkg/autocomplete/model"
	"github.com/wejick/bersih/pkg/autocomplete/repo"
)

// CreateProfile to repo
func (M *MongoRepo) CreateProfile(ctx context.Context, profile model.Profile) (err error) {
	_, err = M.profileCollection.InsertOne(ctx, profile)
	return
}

// GetProfile from repo
func (M *MongoRepo) GetProfile(ctx context.Context) (profileList repo.ProfileList, err error) {
	cur, err := M.profileCollection.Find(context.Background(), nil)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		item := model.Profile{}
		errDecode := cur.Decode(&item)
		if errDecode != nil {
			return profileList, errDecode
		}
	}
	return
}

// UpdateProfile to repo
func (M *MongoRepo) UpdateProfile(ctx context.Context, profile model.Profile) (err error) {
	return
}

// DeleteProfile from repo
func (M *MongoRepo) DeleteProfile(ctx context.Context, profile model.Profile) (err error) {
	return
}
