package service

import (
	autocompleteRepo "github.com/wejick/bersih/pkg/autocomplete/repo"
)

// Service contain the service
type Service struct {
	repo autocompleteRepo.Repo
}

// New instantiate new service
func New(repo autocompleteRepo.Repo) (service *Service) {
	return &Service{
		repo: repo,
	}
}
