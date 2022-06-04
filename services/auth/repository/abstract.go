package repository

import (
	cache "github.com/patrickmn/go-cache"
	"zebrax.id/product/dmaa/core/utils"
	"zebrax.id/product/dmaa/services/auth/repository/postgres"
)

// Repository interface
type Repository interface{}

// AbstractRepository ..
type AbstractRepository struct {
	Repository
	Cache *cache.Cache
}

// NewRepository ...
func NewRepository() *AbstractRepository {
	newRepo := &AbstractRepository{
		Cache: utils.GetCacheHandler(),
	}

	if utils.GetDatasourceInfo() == utils.Postgres {
		newRepo.Repository = &postgres.Repository{
			Connection: utils.GetPostgresHandler(),
		}
	} else {
		newRepo.Repository = nil
	}

	return newRepo
}
