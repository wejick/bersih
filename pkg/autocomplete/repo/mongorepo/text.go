package mongorepo

import (
	"context"

	"github.com/wejick/bersih/pkg/autocomplete/model"
	"github.com/wejick/bersih/pkg/autocomplete/repo"
)

// CreateText to repo
func (M *MongoRepo) CreateText(ctx context.Context, text model.Text) (err error) {
	_, err = M.textCollection.InsertOne(ctx, text)
	return
}

// GetText from repo
func (M *MongoRepo) GetText(ctx context.Context) (textList repo.TextList, err error) {
	cur, err := M.textCollection.Find(context.Background(), nil)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		item := model.Text{}
		errDecode := cur.Decode(&item)
		if errDecode != nil {
			return textList, errDecode
		}
	}
	return
}

// UpdateText to repo
func (M *MongoRepo) UpdateText(ctx context.Context, text model.Text) (err error) {
	return
}

// DeleteText from repo
func (M *MongoRepo) DeleteText(ctx context.Context, text model.Text) (err error) {
	return
}
