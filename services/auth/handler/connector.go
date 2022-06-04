package handler

import (
	"sync"

	"go-microservices/core/cache"
	"go-microservices/services/auth/repository"
	"go-microservices/services/auth/usecase"
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
