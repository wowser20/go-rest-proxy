package types

type GetDummyProductByIDResponse struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Category           string   `json:"category"`
	Price              float32  `json:"price"`
	DiscountPercentage float32  `json:"discountPercentage"`
	Rating             float32  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

type GetDummyProductsResponse struct {
	Products []GetDummyProductByIDResponse `json:"products"`
}
