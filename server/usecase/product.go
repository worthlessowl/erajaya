package usecase

import (
	"fmt"
	"log"

	"github.com/worthlessowl/erajaya/cache"
	"github.com/worthlessowl/erajaya/db"
)

func (usecase *Usecase) InsertProduct(param AddProduct) error {
	resourceParam := convertParamToResource(param)

	err := usecase.db.InsertProduct(resourceParam)
	if err != nil {
		log.Println("Failed to call usecase.db.InsertProduct")
		return err
	}
	return nil
}

func (usecase Usecase) ListProduct(sortBy Filter) ([]Product, error) {
	cacheKey := "key_" + fmt.Sprintf("%d", sortBy)
	resourceProducts, err := usecase.cache.GetProductCache(cacheKey)
	if err != nil {
		log.Println("Failed to call usecase.cache.GetProductCache")
		return nil, err
	}

	if resourceProducts != nil {
		products := convertCacheProductToUsecase(resourceProducts)
		return products, nil
	}

	products, err := usecase.db.ListProduct(db.Filter(sortBy))
	if err != nil {
		log.Println("Failed to call usecase.db.ListProduct")
		return nil, err
	}

	err = usecase.cache.SetProductCache(cacheKey, convertDBProductToCache(products))
	if err != nil {
		log.Println("Failed to call usecase.cache.SetProductCache")
	}

	return convertDBProductToUsecase(products), nil
}

func convertParamToResource(param AddProduct) db.AddProduct {
	return db.AddProduct{
		Name:        param.Name,
		Price:       param.Price,
		Description: param.Description,
		Quantity:    param.Quantity,
	}
}

func convertCacheProductToUsecase(productsResource []cache.Product) []Product {
	products := []Product{}

	for _, prd := range productsResource {
		products = append(products, Product{
			ID:          prd.ID,
			Name:        prd.Name,
			Price:       prd.Price,
			Description: prd.Description,
			Quantity:    prd.Quantity,
			CreateAt:    prd.CreateAt,
		})
	}
	return products
}

func convertDBProductToUsecase(productsResource []db.Product) []Product {
	products := []Product{}

	for _, prd := range productsResource {
		products = append(products, Product{
			ID:          prd.ID,
			Name:        prd.Name,
			Price:       prd.Price,
			Description: prd.Description,
			Quantity:    prd.Quantity,
			CreateAt:    prd.CreateAt,
		})
	}
	return products
}

func convertDBProductToCache(productsResource []db.Product) []cache.Product {
	products := []cache.Product{}

	for _, prd := range productsResource {
		products = append(products, cache.Product{
			ID:          prd.ID,
			Name:        prd.Name,
			Price:       prd.Price,
			Description: prd.Description,
			Quantity:    prd.Quantity,
			CreateAt:    prd.CreateAt,
		})
	}
	return products
}
