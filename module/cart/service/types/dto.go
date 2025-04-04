package types

type GetCartsResult struct {
	Carts []GetCartByIDResult
}

type GetCartByIDResult struct {
	ID              int
	Products        []Product
	Total           float32
	DiscountedTotal float32
	UserID          int
	TotalProducts   int
	TotalQuantity   int
}

type Product struct {
	ID                 int
	Title              string
	Price              float32
	Quantity           int
	Total              float32
	DiscountPercentage float32
	DiscountedTotal    float32
	Thumbnail          string
}
