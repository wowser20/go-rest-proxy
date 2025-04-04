package service

import (
	"go-rest-proxy/module/cart/service/types"
	"go-rest-proxy/utils/api/dummyjson"
)

type CartQueryService struct{}

// GetCarts get all carts
func (service *CartQueryService) GetCarts() ([]types.GetCartsResult, error) {
	res, err := dummyjson.GetDummyCarts()
	if err != nil {
		return nil, err
	}

	carts := []types.GetCartsResult{}

	for _, cart := range res.Carts {
		products := []types.Product{}

		for _, product := range cart.Products {
			products = append(products, types.Product{
				ID:                 product.ID,
				Title:              product.Title,
				Price:              product.Price,
				Quantity:           product.Quantity,
				Total:              product.Total,
				DiscountPercentage: product.DiscountPercentage,
				DiscountedTotal:    product.DiscountedTotal,
				Thumbnail:          product.Thumbnail,
			})
		}

		carts = append(carts, types.GetCartsResult{
			Carts: []types.GetCartByIDResult{
				{
					ID:              cart.ID,
					Products:        products,
					Total:           cart.Total,
					DiscountedTotal: cart.DiscountedTotal,
					UserID:          cart.UserID,
					TotalProducts:   cart.TotalProducts,
					TotalQuantity:   cart.TotalQuantity,
				},
			},
		})
	}

	return carts, nil
}

// GetCartByID get cart by id
func (service *CartQueryService) GetCartByID(ID int) (types.GetCartByIDResult, error) {
	res, err := dummyjson.GetDummyCartByID(ID)
	if err != nil {
		return types.GetCartByIDResult{}, err
	}

	products := []types.Product{}

	for _, product := range res.Products {
		products = append(products, types.Product{
			ID:                 product.ID,
			Title:              product.Title,
			Price:              product.Price,
			Quantity:           product.Quantity,
			Total:              product.Total,
			DiscountPercentage: product.DiscountPercentage,
			DiscountedTotal:    product.DiscountedTotal,
			Thumbnail:          product.Thumbnail,
		})
	}

	cart := types.GetCartByIDResult{
		ID:              res.ID,
		Products:        products,
		Total:           res.Total,
		DiscountedTotal: res.DiscountedTotal,
		UserID:          res.UserID,
		TotalProducts:   res.TotalProducts,
		TotalQuantity:   res.TotalQuantity,
	}

	return cart, nil
}
