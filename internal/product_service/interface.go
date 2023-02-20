package product_service

// ProductServiceInterface Interface to interact with products
type ProductServiceInterface interface {
	GetProducts() ([]ProductStorage, error)
	GetProduct(apiName string) (*ProductStorage, error)
	DeleteProductByID(id string) error
	NewProduct(product *NewProductDTO) error
	UpdateProduct(product *ProductUpdateDTO) error
	GetProductsByCategory(apiName string) ([]ProductStorage, error)
}
