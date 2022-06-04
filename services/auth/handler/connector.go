package handler

import (
	"sync"

	"zebrax.id/product/dmaa/core/cache"
	"zebrax.id/product/dmaa/services/auth/repository"
	"zebrax.id/product/dmaa/services/auth/usecase"
)

var uc *usecase.UseCase
var oneUc sync.Once

// GetUsecase ...
func GetUsecase() *usecase.UseCase {
	oneUc.Do(func() {
		uc = usecase.NewUseCase(
			getRepository(),
			getSession(),
		)
	})

	return uc
}

var (
	repo    *repository.AbstractRepository
	oneRepo sync.Once
)

func getRepository() *repository.AbstractRepository {
	oneRepo.Do(func() {
		repo = repository.NewRepository()
	})

	return repo
}

var (
	sess    cache.ICache
	oneSess sync.Once
)

func getSession() cache.ICache {
	oneSess.Do(func() {
		sess = cache.NewSession()
	})

	return sess
}
