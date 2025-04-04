package types

type GetProductByIDResult struct {
	ID                 int
	Title              string
	Description        string
	Category           string
	Price              float32
	DiscountPercentage float32
	Rating             float32
	Stock              int
	Brand              string
	Thumbnail          string
	Images             []string
}

type GetProductsResult struct {
	Products []GetProductByIDResult
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
