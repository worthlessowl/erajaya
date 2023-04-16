package usecase

import (
	"github.com/worthlessowl/erajaya/cache"
	"github.com/worthlessowl/erajaya/db"
)

type cacheResource interface {
	SetProductCache(key string, products []cache.Product) error

	GetProductCache(key string) ([]cache.Product, error)
}

type databaseResource interface {
	CreateProductTable() error

	InsertProduct(param db.AddProduct) error

	ListProduct(sortBy db.Filter) ([]db.Product, error)
}

type Usecase struct {
	db    databaseResource
	cache cacheResource
}

func New(db databaseResource, redis cacheResource) Usecase {
	return Usecase{db: db, cache: redis}
}
