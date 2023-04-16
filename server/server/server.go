package server

import "github.com/worthlessowl/erajaya/usecase"

type usecaseInterface interface {
	InsertProduct(param usecase.AddProduct) error

	ListProduct(sortBy usecase.Filter) ([]usecase.Product, error)
}

type Server struct {
	uc usecaseInterface
}

func New(uc usecaseInterface) Server {
	return Server{uc: uc}
}
