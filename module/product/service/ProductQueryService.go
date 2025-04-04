package service

import (
	"go-rest-proxy/module/product/service/types"
	"go-rest-proxy/utils/api/dummyjson"
)

type ProductQueryService struct{}

// GetProducts get all products
func (service *ProductQueryService) GetProducts() ([]types.GetProductsResult, error) {
	res, err := dummyjson.GetDummyProducts()
	if err != nil {
		return nil, err
	}

	products := []types.GetProductsResult{}

	for _, product := range res.Products {
		products = append(products, types.GetProductsResult{
			Products: []types.GetProductByIDResult{
				{
					ID:                 product.ID,
					Title:              product.Title,
					Description:        product.Description,
					Category:           product.Category,
					Price:              product.Price,
					DiscountPercentage: product.DiscountPercentage,
					Rating:             product.Rating,
					Stock:              product.Stock,
					Brand:              product.Brand,
					Thumbnail:          product.Thumbnail,
					Images:             product.Images,
				},
			},
		})
	}

	return products, nil
}

// GetProductByID get product by id
func (service *ProductQueryService) GetProductByID(ID int) (types.GetProductByIDResult, error) {
	res, err := dummyjson.GetDummyProductByID(ID)
	if err != nil {
		return types.GetProductByIDResult{}, err
	}

	product := types.GetProductByIDResult{
		ID:                 res.ID,
		Title:              res.Title,
		Description:        res.Description,
		Category:           res.Category,
		Price:              res.Price,
		DiscountPercentage: res.DiscountPercentage,
		Rating:             res.Rating,
		Stock:              res.Stock,
		Brand:              res.Brand,
		Thumbnail:          res.Thumbnail,
		Images:             res.Images,
	}

	return product, nil
}
