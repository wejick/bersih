package mongorepo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	dbName            = "autocomplete"
	textCollection    = "text"
	profileCollection = "profile"
)

// MongoRepo repository backed by mongo db
type MongoRepo struct {
	client            *mongo.Client
	textCollection    *mongo.Collection
	profileCollection *mongo.Collection
}

// New create new mongorepo instance
func New(mongoClient *mongo.Client) *MongoRepo {
	return &MongoRepo{
		client: mongoClient,
	}
}

// Initialize mongorepo
func (M *MongoRepo) Initialize() (err error) {
	M.textCollection = M.client.Database(dbName).Collection(textCollection)
	M.profileCollection = M.client.Database(dbName).Collection(profileCollection)
	return
}

// ProvideRepo provides building repo
func ProvideRepo(mongoClient *mongo.Client) (*MongoRepo, error) {
	repo := New(mongoClient)
	err := repo.Initialize()

	return repo, err
}
