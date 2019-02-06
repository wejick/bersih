package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
	autocompleteMongoRepo "github.com/wejick/bersih/pkg/autocomplete/repo/mongorepo"
	autocompleteService "github.com/wejick/bersih/pkg/autocomplete/service"
)

type appConfig struct {
	appName              string
	AutocompleteMongoURI string
}

func main() {
	config := appConfig{
		appName:              "autocomplete",
		AutocompleteMongoURI: "localhost:27017",
	}
	repo, err := provideRepo(config)
	service := autocompleteService.New(repo)

	return
}

func provideRepo(config appConfig) (repo autocompleteRepo.Repo, err error) {
	mongoClient, err := ProvideMongoClient(config)
	if err != nil {
		return
	}
	repo = autocompleteMongoRepo.New(mongoClient)
	err = repo.Initialize()

	return
}

// ProvideMongoClient provides mongo client
func ProvideMongoClient(config appConfig) (mongoClient *mongo.Client, err error) {
	mongoClient, err = mongo.Connect(nil, config.AutocompleteMongoURI)
	if err != nil {
		return
	}
	err = mongoClient.Ping(nil, nil)
	if err != nil {
		return
	}
	return
}
