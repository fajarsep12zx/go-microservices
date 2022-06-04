package usecase

import (
	"zebrax.id/product/dmaa/services/user/client"
	"zebrax.id/product/dmaa/services/user/repository"
)

// UseCase ...
type UseCase struct {
	repository *repository.AbstractRepository
	client     *client.Client
}

// NewUseCase ...
func NewUseCase(repo *repository.AbstractRepository, cl *client.Client) *UseCase {
	return &UseCase{
		repository: repo,
		client:     cl,
	}
}
