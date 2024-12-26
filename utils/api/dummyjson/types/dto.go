package types

type GetDummyCartsResponse struct {
	Carts []GetDummyCartByIDResponse `json:"carts"`
}

type GetDummyCartByIDResponse struct {
	ID              int       `json:"id"`
	Products        []Product `json:"products"`
	Total           float32   `json:"total"`
	DiscountedTotal float32   `json:"discountedTotal"`
	UserID          int       `json:"userId"`
	TotalProducts   int       `json:"totalProducts"`
	TotalQuantity   int       `json:"totalQuantity"`
}

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

type Product struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Price              float32 `json:"price"`
	Quantity           int     `json:"quantity"`
	Total              float32 `json:"total"`
	DiscountPercentage float32 `json:"discountPercentage"`
	DiscountedTotal    float32 `json:"discountedTotal"`
	Thumbnail          string  `json:"thumbnail"`
}
